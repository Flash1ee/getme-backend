package client

import (
	proto "getme-backend/internal/microservices/auth/delivery/grpc/protobuf"
	"getme-backend/internal/microservices/auth/sessions/models"
)

func ConvertAuthServerRespond(result *proto.Result) models.Result {
	if result == nil {
		return models.Result{}
	}
	res := models.Result{
		UserID: result.UserID,
		UniqID: result.SessionID,
	}
	return res
}

func ConvertAuthServerRespondByToken(result *proto.ResultByToken) models.ResultByToken {
	if result == nil {
		return models.ResultByToken{}
	}
	res := models.ResultByToken{
		TokenID: result.Token,
		UserID:  result.UserID,
	}
	return res
}
