package platform

import (
	"context"
	"sdxx/server/internal/code"
	"sdxx/server/internal/component/snowflake"
	"sdxx/server/internal/dao/mysql/platform/model"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"

	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/utils/xtime"
	"gorm.io/gorm"
)

func (s *Server) Register(ctx context.Context, req *platformv1.RegisterReq) (*platformv1.RegisterResp, error) {
	if isValid := model.AccountAuthChannel(byte(req.Channel)).IsValid(); !isValid {
		return nil, errors.NewError(code.InvalidArgument, "channel")
	}
	if req.DeviceId == "" {
		return nil, errors.NewError(code.InvalidArgument, "device_id")
	}

	var shouldRegister bool
	_, err := s.accountAuthModel.FindOneByChannelIdentifier(ctx, uint64(req.Channel), req.Identifier)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			shouldRegister = true
		} else {
			return nil, err
		}
	}
	if shouldRegister {
		accountId := snowflake.Generate()
		err = s.db.Transaction(func(tx *gorm.DB) error {
			err = s.accountModel.Insert(ctx, tx, &model.Account{
				AccountId:    accountId,
				Mobile:       "",
				DeviceType:   req.DeviceType,
				DeviceId:     req.DeviceId,
				RegisterTime: xtime.Now(),
				RegisterIp:   req.ClientIp,
				Status:       uint64(model.AccountStatusNormal),
			})
			if err != nil {
				return err
			}
			err = s.accountAuthModel.Insert(ctx, tx, &model.AccountAuth{
				AccountId:  accountId,
				Channel:    uint64(req.Channel),
				Identifier: req.Identifier,
				Status:     uint64(model.AccountAuthStatusNormal),
			})
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return &platformv1.RegisterResp{
		IsNewlyRegistered: shouldRegister,
	}, nil
}
