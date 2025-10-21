package platform

import (
	"context"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"

	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/jwt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) RefreshToken(ctx context.Context, req *platformv1.RefreshTokenReq) (*platformv1.RefreshTokenResp, error) {
	token, err := s.jwt.RefreshToken(req.AccessToken)
	if err != nil {
		switch {
		case jwt.IsExpiredToken(err):
			return nil, errors.NewError(commonv1.TokenExpired)
		case jwt.IsMissingToken(err):
			return nil, errors.NewError(commonv1.TokenMissing)
		default:
			return nil, errors.NewError(commonv1.TokenInvalid)
		}
	}
	return &platformv1.RefreshTokenResp{
		AccessToken:  token.Token,
		AccessExpire: timestamppb.New(token.ExpiredAt),
		RefreshAfter: timestamppb.New(token.RefreshAt),
	}, nil
}
