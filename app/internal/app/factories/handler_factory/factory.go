package handler_factory

import (
	"getme-backend/internal/app"
	user_create_handler "getme-backend/internal/app/user/delivery/http/handlers/user_create"
	user_profile_handler "getme-backend/internal/app/user/delivery/http/handlers/user_profile"

	"github.com/sirupsen/logrus"
)

const (
	PROFILE = iota
	USER_CREATE
)

type HandlerFactory struct {
	repositoryFactory RepositoryFactory
	logger            *logrus.Logger
	urlHandler        *map[string]app.Handler
}

func NewFactory(logger *logrus.Logger, repositoryFactory RepositoryFactory) *HandlerFactory {
	return &HandlerFactory{
		repositoryFactory: repositoryFactory,
		logger:            logger,
	}
}

func (f *HandlerFactory) initAllHandlers() map[int]app.Handler {
	return map[int]app.Handler{
		PROFILE:     user_profile_handler.NewUserProfileHandler(f.logger, f.repositoryFactory.GetUserRepository()),
		USER_CREATE: user_create_handler.NewUserCreateHandler(f.logger, f.repositoryFactory.GetUserRepository()),
	}
}

func (f *HandlerFactory) GetHandleUrls() *map[string]app.Handler {
	if f.urlHandler != nil {
		return f.urlHandler
	}

	hs := f.initAllHandlers()
	f.urlHandler = &map[string]app.Handler{
		//=============user==============//
		"/user/<nickname>/profile": hs[PROFILE],
		"/user/<nickname>/create":  hs[USER_CREATE],
	}
	return f.urlHandler
}
