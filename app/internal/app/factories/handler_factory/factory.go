package handler_factory

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"getme-backend/internal/app"
	"getme-backend/internal/app/token/delivery/http/handlers/token_handler"
	"getme-backend/internal/app/user/delivery/http/handlers/user_auth/simple"
	user_auth_handler "getme-backend/internal/app/user/delivery/http/handlers/user_auth/telegram"
	user_auth_check_handler "getme-backend/internal/app/user/delivery/http/handlers/user_auth_check/telegram"
	"getme-backend/internal/app/user/delivery/http/handlers/user_logout"
	"getme-backend/internal/app/user/delivery/http/handlers/user_register"
	"getme-backend/internal/microservices/auth/delivery/grpc/client"
)

const (
	AUTH_CHECK = iota
	AUTH
	AUTH_TOKEN
	LOGIN
	LOGOUT
	REGISTER
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
	tokenUsecase := f.usecaseFactory.GetTokenUsecase()

	sClient := client.NewSessionClient(f.sessionClientConn)
	return map[int]app.Handler{
		AUTH_CHECK: user_auth_check_handler.NewUserAuthCheckHandler(f.logger, ucUsecase, sClient),
		AUTH:       user_auth_handler.NewUserAuthHandler(f.logger, ucUsecase, sClient, tokenUsecase),
		AUTH_TOKEN: token_handler.NewTokenHandler(f.logger, tokenUsecase, sClient),
		LOGOUT:     user_logout.NewLogoutHandler(f.logger, sClient),
		LOGIN:      user_simple_auth.NewLoginHandler(f.logger, sClient, ucUsecase),
		REGISTER:   user_register.NewRegisterHandler(f.logger, sClient, ucUsecase),
	}
}

func (f *HandlerFactory) GetHandleUrls() *map[string]app.Handler {
	if f.urlHandler != nil {
		return f.urlHandler
	}

	hs := f.initAllHandlers()
	f.urlHandler = &map[string]app.Handler{
		//=============user==============//
		"/auth/telegram/callback": hs[AUTH],

		//"/login":         hs[AUTH],
		//"/login/confirm": hs[AUTH],
		//
		//"/register/telegram": hs[AUTH],
		//"/register/confirm":  hs[AUTH],

		"/auth/token":           hs[AUTH_TOKEN],
		"/auth/telegram/check":  hs[AUTH_CHECK],
		"/logout":               hs[LOGOUT],
		"/auth/simple/login":    hs[LOGIN],
		"/auth/simple/register": hs[REGISTER],

		//"/user/<nickname>/create":  hs[USER_CREATE],
	}
	return f.urlHandler
}
