package client

import (
	"context"

	"getme-backend/internal/microservices/auth/sessions/models"

	proto "getme-backend/internal/microservices/auth/delivery/grpc/protobuf"

	"google.golang.org/grpc"
)

type SessionClient struct {
	sessionClient proto.AuthCheckerClient
}

func NewSessionClient(con *grpc.ClientConn) *SessionClient {
	client := proto.NewAuthCheckerClient(con)
	return &SessionClient{
		sessionClient: client,
	}
}

// Check Errors:
//		Status 401 "not authorized user"
func (c *SessionClient) Check(ctx context.Context, sessionID string) (models.Result, error) {
	protoSessionID := &proto.SessionID{ID: sessionID}
	res, err := c.sessionClient.Check(ctx, protoSessionID)
	if err != nil {
		return models.Result{}, err
	}
	return ConvertAuthServerRespond(res), err
}
func (c *SessionClient) Create(ctx context.Context, userID int64) (models.Result, error) {
	protoUserID := &proto.UserID{
		ID: userID,
	}
	res, err := c.sessionClient.Create(ctx, protoUserID)
	if err != nil {
		return models.Result{}, err
	}
	return ConvertAuthServerRespond(res), nil
}
func (c *SessionClient) CreateByToken(ctx context.Context, tokenID string, userID int64) (models.ResultByToken, error) {
	protoTokenData := &proto.TokenUserData{
		Token: &proto.TokenID{
			ID: tokenID,
		},
		User: &proto.UserID{
			ID: userID,
		},
	}
	res, err := c.sessionClient.CreateByToken(ctx, protoTokenData)
	if err != nil {
		return models.ResultByToken{}, nil
	}
	return ConvertAuthServerRespondByToken(res), nil
}

func (c *SessionClient) Delete(ctx context.Context, sessionID string) error {
	protoSessionID := &proto.SessionID{
		ID: sessionID,
	}
	_, err := c.sessionClient.Delete(ctx, protoSessionID)
	if err != nil {
		return err
	}

	return nil
}
func (c *SessionClient) CheckWithDelete(ctx context.Context, tokenID string) (models.ResultByToken, error) {
	protoSessionID := &proto.TokenID{ID: tokenID}
	res, err := c.sessionClient.CheckWithDelete(ctx, protoSessionID)
	if err != nil {
		return models.ResultByToken{}, err
	}
	return ConvertAuthServerRespondByToken(res), err
}
