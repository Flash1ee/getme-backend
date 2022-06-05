package server

import (
	"context"

	proto "getme-backend/internal/microservices/auth/delivery/grpc/protobuf"
)

type AuthCheckerServer interface {
	Check(context.Context, *proto.SessionID) (*proto.Result, error)
	Create(context.Context, *proto.UserID) (*proto.Result, error)
	Delete(context.Context, *proto.SessionID) (*proto.Nothing, error)
	CreateByToken(context.Context, *proto.TokenUserData) (*proto.ResultByToken, error)
	CheckWithDelete(context.Context, *proto.TokenID) (*proto.ResultByToken, error)
}
