package platform

import (
	"context"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"

	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/utils/xconv"
)

func (s *Server) ValidateToken(ctx context.Context, req *platformv1.ValidateTokenReq) (*platformv1.ValidateTokenResp, error) {
	identity, err := s.jwt.ExtractIdentity(req.AccessToken)
	if err != nil {
		return nil, errors.NewError(err, commonv1.Unauthorized)
	}

	accountId := xconv.Uint64(identity)
	if accountId <= 0 {
		return nil, errors.NewError(err, commonv1.Unauthorized)
	}

	return &platformv1.ValidateTokenResp{AccountId: accountId}, nil
}
