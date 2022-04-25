package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"

	"getme-backend/internal"
	"getme-backend/internal/app/factories/handler_factory"
	"getme-backend/internal/app/factories/repository_factory"
	"getme-backend/internal/app/factories/usecase_factory"
	"getme-backend/internal/app/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	"getme-backend/internal/pkg/server"
	"getme-backend/internal/pkg/utilits"
)

type Server struct {
	server.BaseServer
}

func New(config *internal.Config, connections utilits.ExpectedConnections, logger *log.Logger) *Server {
	return &Server{
		BaseServer: *server.NewBaseServer(config, connections, logger),
	}
}

func (s *Server) Start(config *internal.Config) error {
	if err := s.Check(); err != nil {
		return err
	}

	router := echo.New()
	router.Renderer = echo_adapter.NewRenderer("/Users/dvvarin/TP/getme-backend/app/template/*", false)
	utilityMiddleware := middleware.NewUtilitiesMiddleware(s.Logger)

	//router.Get("/debug/pprof/<profile>", handler_interfaces.FastHTTPFunc(pprofhandler.PprofHandler).ServeHTTP)

	routerApi := router.Group("/api")

	repositoryFactory := repository_factory.NewRepositoryFactory(s.Logger, s.Connections)
	usecaseFactory := usecase_factory.NewUsecaseFactory(s.Logger, repositoryFactory, config.TgAuth)
	factory := handler_factory.NewFactory(s.Logger, usecaseFactory)

	hs := factory.GetHandleUrls()

	routerApi.Use(
		echo.WrapMiddleware(utilityMiddleware.CheckPanic),
		echo.WrapMiddleware(utilityMiddleware.UpgradeLogger))

	for apiUrl, h := range *hs {
		h.Connect(routerApi, apiUrl)
	}

	s.Logger.Info("start http server")

	router.GET(
		"/login", func(c echo.Context) error {
			c.Render(http.StatusOK, "login.tmpl", map[string]interface{}{})
			return nil
		})

	return router.Start(config.BindAddr)
}
