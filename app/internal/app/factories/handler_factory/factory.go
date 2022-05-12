package handler_factory

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"getme-backend/internal/app"
	"getme-backend/internal/app/token/delivery/http/handlers/token_handler"
	user_auth_handler "getme-backend/internal/app/user/delivery/http/handlers/user_auth"
	user_auth_check_handler "getme-backend/internal/app/user/delivery/http/handlers/user_auth_check"
	"getme-backend/internal/microservices/auth/delivery/grpc/client"
)

const (
	AUTH_CHECK = iota
	AUTH
	TOKEN_AUTH
)

type HandlerFactory struct {
	usecaseFactory    UsecaseFactory
	sessionClientConn *grpc.ClientConn
	logger            *logrus.Logger
	urlHandler        *map[string]app.Handler
}

func NewFactory(logger *logrus.Logger, sessionClientConn *grpc.ClientConn, ucFactory UsecaseFactory) *HandlerFactory {
	return &HandlerFactory{
		usecaseFactory:    ucFactory,
		logger:            logger,
		sessionClientConn: sessionClientConn,
	}
}

func (f *HandlerFactory) initAllHandlers() map[int]app.Handler {
	ucUsecase := f.usecaseFactory.GetUserUsecase()

	sClient := client.NewSessionClient(f.sessionClientConn)
	return map[int]app.Handler{
		AUTH_CHECK: user_auth_check_handler.NewUserAuthCheckHandler(f.logger, ucUsecase, sClient),
		AUTH:       user_auth_handler.NewUserAuthHandler(f.logger, ucUsecase, sClient),
		TOKEN_AUTH: token_handler.NewTokenHandler(f.logger, nil, sClient),
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
		"/auth/check": hs[AUTH_CHECK],
		"/auth":       hs[AUTH],
		//"/user/<nickname>/create":  hs[USER_CREATE],
	}
	return f.urlHandler
}
