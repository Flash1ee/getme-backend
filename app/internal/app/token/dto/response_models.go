package dto

import (
	"getme-backend/internal/app/token/entities"
)

//go:generate easyjson -disallow_unknown_fields response_models.go

//easyjson:json
type TokenResponse struct {
	Token string `json:"token"`
}

func ToTokenResponse(token *entities.Token) *TokenResponse {
	return &TokenResponse{
		Token: token.Token,
	}
}
