package usecase_factory

import (
	"github.com/sirupsen/logrus"

	"getme-backend/internal"
	"getme-backend/internal/app/auth/services/telegram-checker"
	authUs "getme-backend/internal/app/auth/usecase"
	"getme-backend/internal/app/auth/usecase/auth_usecase"
	skillUs "getme-backend/internal/app/skill/usecase"
	"getme-backend/internal/app/skill/usecase/skill_usecase"
	tokenUs "getme-backend/internal/app/token/usecase"
	"getme-backend/internal/app/token/usecase/token_usecase"
	userUs "getme-backend/internal/app/user/usecase"
	"getme-backend/internal/app/user/usecase/user_usecase"
)

type UsecaseFactory struct {
	log               logrus.Logger
	repositoryFactory RepositoryFactory
	userUsecase       userUs.Usecase
	tokenUsecase      tokenUs.Usecase
	authUsecase       authUs.Usecase
	skillUsecase      skillUs.Usecase

	authChecker *telegram_checker.TelegramChecker
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
func (f *UsecaseFactory) GetAuthUsecase() authUs.Usecase {
	if f.authUsecase == nil {
		f.authUsecase = auth_usecase.NewAuthUsecase(f.repositoryFactory.GetAuthRepository(),
			f.repositoryFactory.GetUserRepository(), f.authChecker)
	}
	return f.authUsecase
}
func (f *UsecaseFactory) GetSkillUsecase() skillUs.Usecase {
	if f.skillUsecase == nil {
		f.skillUsecase = skill_usecase.NewSkillUsecase(f.repositoryFactory.GetSkillRepository())
	}

	return f.skillUsecase
}
