package repository_factory

import (
	"github.com/sirupsen/logrus"

	authRepo "getme-backend/internal/app/auth/repository"
	offerRepo "getme-backend/internal/app/offer/repository"
	offerPostgresRepo "getme-backend/internal/app/offer/repository/postgresql"
	planRepo "getme-backend/internal/app/plans/repository"
	planPostgresRepo "getme-backend/internal/app/plans/repository/postgresql"
	skillRepo "getme-backend/internal/app/skill/repository"
	taskRepo "getme-backend/internal/app/task/repository"
	taskPostgresRepo "getme-backend/internal/app/task/repository/postgresql"

	skillPostgresRepo "getme-backend/internal/app/skill/repository/postgresql"

	authPostgresRepo "getme-backend/internal/app/auth/repository/postgresql"
	tokenRepo "getme-backend/internal/app/token/repository"
	token_jwt_repository "getme-backend/internal/app/token/repository/jwt"
	token_redis_repository "getme-backend/internal/app/token/repository/redis"
	userRepo "getme-backend/internal/app/user/repository"
	repository_postgresql "getme-backend/internal/app/user/repository/postgresql"
	"getme-backend/internal/pkg/utilits"
)

type RepositoryFactory struct {
	expectedConnections utilits.ExpectedConnections
	logger              *logrus.Logger
	userRepository      userRepo.Repository
	authRepository      authRepo.Repository
	tokenRepository     tokenRepo.Repository
	tokenJWTRepository  tokenRepo.RepositoryJWT
	skillRepository     skillRepo.Repository
	offerRepository     offerRepo.Repository
	planRepository      planRepo.Repository
	taskRepository      taskRepo.Repository
}

func NewRepositoryFactory(logger *logrus.Logger, expectedConnections utilits.ExpectedConnections) *RepositoryFactory {
	return &RepositoryFactory{
		expectedConnections: expectedConnections,
		logger:              logger,
	}
}

func (f *RepositoryFactory) GetUserRepository() userRepo.Repository {
	if f.userRepository == nil {
		f.userRepository = repository_postgresql.NewUserRepository(f.expectedConnections.SqlConnection)
	}
	return f.userRepository
}
func (f *RepositoryFactory) GetTokenRepository() tokenRepo.Repository {
	if f.tokenRepository == nil {
		f.tokenRepository = token_redis_repository.NewTokenRepository(f.expectedConnections.UtilsRedisPool)
	}
	return f.tokenRepository
}
func (f *RepositoryFactory) GetTokenJWTRepository() tokenRepo.RepositoryJWT {
	if f.tokenJWTRepository == nil {
		f.tokenJWTRepository = token_jwt_repository.NewJwtRepository()
	}
	return f.tokenJWTRepository
}
func (f *RepositoryFactory) GetAuthRepository() authRepo.Repository {
	if f.authRepository == nil {
		f.authRepository = authPostgresRepo.NewAuthRepository(f.expectedConnections.SqlConnection)
	}
	return f.authRepository
}

func (f *RepositoryFactory) GetSkillRepository() skillRepo.Repository {
	if f.skillRepository == nil {
		f.skillRepository = skillPostgresRepo.NewSkillRepository(f.expectedConnections.SqlConnection)
	}
	return f.skillRepository
}

func (f *RepositoryFactory) GetOfferRepository() offerRepo.Repository {
	if f.offerRepository == nil {
		f.offerRepository = offerPostgresRepo.NewOfferRepository(f.expectedConnections.SqlConnection)
	}
	return f.offerRepository
}

func (f *RepositoryFactory) GetPlanRepository() planRepo.Repository {
	if f.planRepository == nil {
		f.planRepository = planPostgresRepo.NewPlanRepository(f.expectedConnections.SqlConnection)
	}
	return f.planRepository
}
func (f *RepositoryFactory) GetTaskRepository() taskRepo.Repository {
	if f.taskRepository == nil {
		f.taskRepository = taskPostgresRepo.NewTaskRepository(f.expectedConnections.SqlConnection)
	}
	return f.taskRepository
}
