package platform

import (
	"context"
	"database/sql"
	"sdxx/server/internal/code"
	"sdxx/server/internal/dao/mysql/platform/model"
	"sdxx/server/internal/types"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"

	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/session"
	"github.com/dobyte/due/v2/utils/xconv"
	"github.com/dobyte/due/v2/utils/xtime"
	"github.com/dobyte/jwt"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (s *Server) doCheckRemoteLogin(ctx context.Context, accountId uint64) error {
	if err := s.jwt.DestroyIdentity(accountId); err != nil {
		log.Errorf("destory user's identity failed: %v", err)
	}

	accountPlayers, err := s.accountPlayerModel.FindByAccountId(ctx, int64(accountId))
	if err != nil {
		log.Errorf("find account players by account id failed: %v", err)
		return errors.NewError(commonv1.InternalError)
	}

	for _, accountPlayer := range accountPlayers {
		gid, err := s.proxy.LocateGate(ctx, int64(accountPlayer.PlayerId))
		if err != nil && !errors.Is(err, errors.ErrNotFoundUserLocation) {
			log.Errorf("locate user's gate failed: %v", err)
			return errors.NewError(commonv1.InternalError)
		}
		if gid == "" {
			continue
		}

		err = s.proxy.Push(ctx, &cluster.PushArgs{
			Kind:    session.User,
			Target:  int64(accountId),
			Message: &cluster.Message{Route: int32(commonv1.Route_ROUTE_REMOTE_LOGIN_NOTIFY)},
		})
		if err != nil {
			log.Errorf("push remote login notify failed: %v", err)
			continue
		}
	}

	return nil
}

func (s *Server) Login(ctx context.Context, req *platformv1.LoginReq) (*platformv1.LoginResp, error) {
	if req.DeviceId == "" {
		return nil, errors.NewError(code.InvalidArgument, "device_id")
	}

	channel, err := types.ParseChannel(int(req.Channel))
	if err != nil {
		return nil, errors.NewError(code.InvalidArgument, "channel")
	}

	accountAuth, err := s.accountAuthModel.FindOneByChannelIdentifier(ctx, uint64(channel), req.Identifier)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.NewError(commonv1.AccountNotRegistered)
		} else {
			return nil, err
		}
	}
	account, err := s.accountModel.FindOneByAccountId(ctx, accountAuth.AccountId)
	if err != nil {
		return nil, errors.NewError(commonv1.AccountDataAbnormal)
	}
	if account.Status != uint64(model.AccountAuthStatusNormal) {
		return nil, errors.NewError(commonv1.AccountDisabled)
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		// 更新用户最新登录信息
		err = s.accountModel.UpdateById(ctx, tx, account.Id, &model.Account{
			Id:              account.Id,
			LastLoginTime:   sql.NullTime{Time: xtime.Now(), Valid: true},
			LastLoginIp:     req.ClientIp,
			LastLoginDevice: req.DeviceId,
		})
		if err != nil {
			return err
		}
		err = s.accountAuthModel.UpdateById(ctx, tx, accountAuth.Id, &model.AccountAuth{
			Id:           accountAuth.Id,
			LastUsedTime: sql.NullTime{Time: xtime.Now(), Valid: true},
		})
		if err != nil {
			return err
		}

		// 检测异地登录
		if err = s.doCheckRemoteLogin(ctx, account.AccountId); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// 生成Token
	token, err := s.jwt.GenerateToken(jwt.Payload{
		s.jwt.IdentityKey(): xconv.String(account.AccountId),
	})
	if err != nil {
		return nil, errors.NewError(code.InternalError)
	}

	return &platformv1.LoginResp{
		AccessToken:  token.Token,
		AccessExpire: timestamppb.New(token.ExpiredAt),
		RefreshAfter: timestamppb.New(token.RefreshAt),
	}, nil
}
