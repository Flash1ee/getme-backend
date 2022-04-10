package server

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"

	"getme-backend/internal"
	"getme-backend/internal/app"
	"getme-backend/internal/app/factories/handler_factory"
	"getme-backend/internal/app/factories/repository_factory"
	"getme-backend/internal/app/middleware"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	config      *internal.Config
	logger      *log.Logger
	connections app.ExpectedConnections
}

func New(config *internal.Config, connections app.ExpectedConnections, logger *log.Logger) *Server {
	return &Server{
		config:      config,
		logger:      logger,
		connections: connections,
	}
}

func (s *Server) checkConnection() error {
	if err := s.connections.SqlConnection.Ping(); err != nil {
		return fmt.Errorf("Can't check connection to sql with error %v ", err)
	}

	s.logger.Info("Success check connection to sql db")

	return nil
}

func (s *Server) Start(config *internal.Config) error {
	if err := s.checkConnection(); err != nil {
		return err
	}

	router := routing.New()
	//router.Get("/debug/pprof/<profile>", handler_interfaces.FastHTTPFunc(pprofhandler.PprofHandler).ServeHTTP)

	routerApi := router.Group("/api")

	repositoryFactory := repository_factory.NewRepositoryFactory(s.logger, s.connections)
	factory := handler_factory.NewFactory(s.logger, repositoryFactory)
	hs := factory.GetHandleUrls()

	utilityMiddleware := middleware.NewUtilitiesMiddleware(s.logger)
	routerApi.Use(utilityMiddleware.UpgradeLogger().ServeHTTP, utilityMiddleware.CheckPanic().ServeHTTP)

	for apiUrl, h := range *hs {
		h.Connect(routerApi.Connect(apiUrl))
	}

	s.logger.Info("start no http server")
	return fasthttp.ListenAndServe(config.BindAddr, router.HandleRequest)
}
