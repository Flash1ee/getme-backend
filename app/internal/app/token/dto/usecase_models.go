package dto

import (
	"time"

	"getme-backend/internal/app/token/entities"
)

type TokenUsecase struct {
	Token string `json:"token"`
}

func (req *TokenUsecase) ToTokenEntity() *entities.Token {
	return &entities.Token{
		Token: req.Token,
	}
}

type TokenSourcesUsecase struct {
	IdentifierData string
}

func (req *TokenSourcesUsecase) ToTokenSourcesEntity(exp time.Duration) *entities.TokenSources {
	return &entities.TokenSources{
		IdentifierData: req.IdentifierData,
		ExpiredTime:    time.Now().Add(exp),
	}
}
