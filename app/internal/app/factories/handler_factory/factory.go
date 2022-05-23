package handler_factory

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"getme-backend/internal/app"
	"getme-backend/internal/app/auth/delivery/http/handlers/login/simple_auth_handler"
	"getme-backend/internal/app/auth/delivery/http/handlers/login/telegram_auth_handler"
	logout_handler "getme-backend/internal/app/auth/delivery/http/handlers/logout"
	"getme-backend/internal/app/auth/delivery/http/handlers/register/simple_register_handler"
	"getme-backend/internal/app/auth/delivery/http/handlers/register/telegram_register_handler"
	skills_info_handler "getme-backend/internal/app/skill/delivery/http/skills/info_handler"
	"getme-backend/internal/app/token/delivery/http/handlers/token_handler"
	user_profile_handler "getme-backend/internal/app/user/delivery/http/handlers/profile"
	"getme-backend/internal/microservices/auth/delivery/grpc/client"
)

const (
	AUTH_TG = iota
	AUTH_SIMPLE
	AUTH_TOKEN
	REGISTER_TG
	REGISTER_SIMPLE
	LOGOUT
	SKILL_INFO
	PROFILE
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
	authUsecase := f.usecaseFactory.GetAuthUsecase()
	skillUsecase := f.usecaseFactory.GetSkillUsecase()

	sClient := client.NewSessionClient(f.sessionClientConn)
	return map[int]app.Handler{
		LOGOUT:          logout_handler.NewLogoutHandler(f.logger, sClient),
		AUTH_TG:         telegram_auth_handler.NewAuthHandler(f.logger, sClient, tokenUsecase),
		REGISTER_TG:     telegram_register_handler.NewRegisterHandler(f.logger, sClient, authUsecase, ucUsecase),
		AUTH_SIMPLE:     simple_auth_handler.NewAuthHandler(f.logger, sClient, authUsecase),
		REGISTER_SIMPLE: simple_register_handler.NewRegisterHandler(f.logger, sClient, ucUsecase, authUsecase),
		AUTH_TOKEN:      token_handler.NewTokenHandler(f.logger, tokenUsecase, sClient),
		SKILL_INFO:      skills_info_handler.NewInfoHandler(f.logger, sClient, skillUsecase),
		PROFILE:         user_profile_handler.NewProfileHandler(f.logger, sClient, ucUsecase),
	}
}

func (f *HandlerFactory) GetHandleUrls() *map[string]app.Handler {
	if f.urlHandler != nil {
		return f.urlHandler
	}

	hs := f.initAllHandlers()
	f.urlHandler = &map[string]app.Handler{
		//=============auth==============//
		"/logout":                 hs[LOGOUT],
		"/auth/telegram/login":    hs[AUTH_TG],
		"/auth/simple/login":      hs[AUTH_SIMPLE],
		"/auth/telegram/register": hs[REGISTER_TG],
		"/auth/simple/register":   hs[REGISTER_SIMPLE],
		"/auth/token":             hs[AUTH_TOKEN],
		//=============skills=============//
		"/skills": hs[SKILL_INFO],
		//=============user=============//
		"/profile": hs[PROFILE],
	}
	return f.urlHandler
}
