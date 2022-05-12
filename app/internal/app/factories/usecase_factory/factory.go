package usecase_factory

import (
	"github.com/sirupsen/logrus"

	"getme-backend/internal"
	tokenUs "getme-backend/internal/app/token/usecase"
	"getme-backend/internal/app/token/usecase/token_usecase"
	telegram_checker "getme-backend/internal/app/user/services/telegram-checker"
	userUs "getme-backend/internal/app/user/usecase"
	"getme-backend/internal/app/user/usecase/user_usecase"
)

type UsecaseFactory struct {
	log               logrus.Logger
	repositoryFactory RepositoryFactory
	userUsecase       userUs.Usecase
	tokenUsecase      tokenUs.Usecase
	authChecker       *telegram_checker.TelegramChecker
}

func NewUsecaseFactory(log *logrus.Logger, repositoryFactory RepositoryFactory, authConf internal.TelegramAuth) *UsecaseFactory {
	authChecker := telegram_checker.NewTelegramChecker(log, authConf)
	return &UsecaseFactory{
		repositoryFactory: repositoryFactory,
		authChecker:       authChecker,
	}
}

func (f *UsecaseFactory) GetUserUsecase() userUs.Usecase {
	if f.userUsecase == nil {
		f.userUsecase = user_usecase.NewUserUsecase(f.repositoryFactory.GetUserRepository(), f.authChecker)
	}
	return f.userUsecase
}
func (f *UsecaseFactory) GetTokenUsecase() tokenUs.Usecase {
	if f.tokenUsecase == nil {
		f.tokenUsecase = token_usecase.NewTokenUsecase(f.repositoryFactory.GetTokenRepository(),
			f.repositoryFactory.GetTokenJWTRepository())
	}
	return f.tokenUsecase
}
