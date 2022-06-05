package usecase_factory

import (
	"github.com/sirupsen/logrus"

	taskUs "getme-backend/internal/app/task/usecase"
	task_usecase "getme-backend/internal/app/task/usecase/usecase"

	"getme-backend/internal"
	"getme-backend/internal/app/auth/services/telegram-checker"
	authUs "getme-backend/internal/app/auth/usecase"
	"getme-backend/internal/app/auth/usecase/auth_usecase"
	offerUs "getme-backend/internal/app/offer/usecase"
	offer_usecase "getme-backend/internal/app/offer/usecase/usecase"
	plansUs "getme-backend/internal/app/plans/usecase"

	plans_usecase "getme-backend/internal/app/plans/usecase/usecase"
	skillUs "getme-backend/internal/app/skill/usecase"
	"getme-backend/internal/app/skill/usecase/skill_usecase"
	tokenUs "getme-backend/internal/app/token/usecase"
	"getme-backend/internal/app/token/usecase/token_usecase"
	userUs "getme-backend/internal/app/user/usecase"

	"getme-backend/internal/app/user/usecase/user_usecase"
)

type UsecaseFactory struct {
	repositoryFactory RepositoryFactory
	userUsecase       userUs.Usecase
	tokenUsecase      tokenUs.Usecase
	authUsecase       authUs.Usecase
	skillUsecase      skillUs.Usecase
	offersUsecase     offerUs.Usecase
	plansUsecase      plansUs.Usecase
	taskUsecase       taskUs.Usecase

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
		f.skillUsecase = skill_usecase.NewSkillUsecase(f.repositoryFactory.GetSkillRepository(), f.repositoryFactory.GetUserRepository())
	}
	return f.skillUsecase
}
func (f *UsecaseFactory) GetOfferUsecase() offerUs.Usecase {
	if f.offersUsecase == nil {
		f.offersUsecase = offer_usecase.NewOfferUsecase(f.repositoryFactory.GetOfferRepository(), f.repositoryFactory.GetUserRepository(), f.repositoryFactory.GetSkillRepository(), f.repositoryFactory.GetPlanRepository())

	}
	return f.offersUsecase
}

func (f *UsecaseFactory) GetPlansUsecase() plansUs.Usecase {
	if f.plansUsecase == nil {
		f.plansUsecase = plans_usecase.NewPlanUsecase(f.repositoryFactory.GetPlanRepository())

	}
	return f.plansUsecase
}
func (f *UsecaseFactory) GetTaskUsecase() taskUs.Usecase {
	if f.taskUsecase == nil {
		f.taskUsecase = task_usecase.NewTaskUsecase(f.repositoryFactory.GetTaskRepository(), f.repositoryFactory.GetPlanRepository())

	}
	return f.taskUsecase

}
