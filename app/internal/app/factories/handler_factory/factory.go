package handler_factory

import (
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app"
	user_auth_handler "getme-backend/internal/app/user/delivery/http/handlers/user_auth"
)

const (
	AUTH = iota
	PROFILE
	USER_CREATE
)

type HandlerFactory struct {
	usecaseFactory UsecaseFactory
	logger         *logrus.Logger
	urlHandler     *map[string]app.Handler
}

func NewFactory(logger *logrus.Logger, ucFactory UsecaseFactory) *HandlerFactory {
	return &HandlerFactory{
		usecaseFactory: ucFactory,
		logger:         logger,
	}
}

func (f *HandlerFactory) initAllHandlers() map[int]app.Handler {
	ucUsecase := f.usecaseFactory.GetUserUsecase()
	return map[int]app.Handler{
		AUTH: user_auth_handler.NewUserAuthHandler(f.logger, ucUsecase),
		//PROFILE: user_profile_handler.NewUserProfileHandler(f.logger, f.repositoryFactory.GetUserRepository()),
		//USER_CREATE: user_create_handler.NewUserCreateHandler(f.logger, f.repositoryFactory.GetUserRepository()),
	}
}

func (f *HandlerFactory) GetHandleUrls() *map[string]app.Handler {
	if f.urlHandler != nil {
		return f.urlHandler
	}

	hs := f.initAllHandlers()
	f.urlHandler = &map[string]app.Handler{
		//=============user==============//
		"/auth": hs[AUTH],
		//"/user/<nickname>/create":  hs[USER_CREATE],
	}
	return f.urlHandler
}
