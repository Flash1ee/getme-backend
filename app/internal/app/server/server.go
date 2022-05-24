package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"

	_ "getme-backend/statik"

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

	pwd, err := os.Getwd()
	if err != nil {
		s.Logger.Fatal(err)
	}

	router := echo.New()
	//router.Use(echoMW.CSRFWithConfig(echoMW.CSRFConfig{}))
	router.Renderer = echo_adapter.NewRenderer(pwd+"/template/*", false)
	utilityMiddleware := middleware.NewUtilitiesMiddleware(s.Logger)

	routerApi := router.Group("/api/v1")

	repositoryFactory := repository_factory.NewRepositoryFactory(s.Logger, s.Connections)
	usecaseFactory := usecase_factory.NewUsecaseFactory(s.Logger, repositoryFactory, config.TgAuth)
	factory := handler_factory.NewFactory(s.Logger, s.Connections.SessionGrpcConnection, usecaseFactory)

	hs := factory.GetHandleUrls()

	routerApi.Use(
		echo.WrapMiddleware(utilityMiddleware.CheckPanic),
		echo.WrapMiddleware(utilityMiddleware.UpgradeLogger))

	for apiUrl, h := range *hs {
		h.Connect(routerApi, apiUrl)
	}

	swaggerFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(swaggerFS)
	routerApi.GET("/swagger/editor", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/api/v1/editor/swagger-editor/")
	})

	routerApi.GET("/swagger/ui", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/api/v1/editor/swagger-ui/")
	})

	routerApi.GET("/editor/swagger-editor/*", echo.WrapHandler(http.StripPrefix("/api/v1/editor/", staticServer)))
	routerApi.GET("/editor/swagger-ui/*", echo.WrapHandler(http.StripPrefix("/api/v1/editor/", staticServer)))

	s.Logger.Info("start http server")

	router.GET(
		"/login", func(c echo.Context) error {
			err = c.Render(http.StatusOK, "login.tmpl", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		})

	return router.Start(config.BindAddr)
}
