package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"

	"getme-backend/internal/app/auth/dto"
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
		"/login", func(ctx echo.Context) error {
			req := &dto.AuthRequest{}
			binder := echo.QueryParamsBinder(ctx)

			errs := binder.String("token", &req.Token).
				BindErrors()

			if errs != nil {
				for _, err := range errs {
					bErr := err.(*echo.BindingError)
					s.Logger.Errorf("/login error parse query params:field = %v value = %v\n", bErr.Field, bErr.Values)
					return err
				}
				ctx.Response().WriteHeader(http.StatusBadRequest)
			}

			err = ctx.Render(http.StatusOK, "login.tmpl", req)
			if err != nil {
				return err
			}
			return nil
		})

	return router.Start(config.BindAddr)
}
