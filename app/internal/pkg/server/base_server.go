package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"getme-backend/internal"
	"getme-backend/internal/pkg/utilits"
)

type BaseServer struct {
	Config      *internal.Config
	Logger      *log.Logger
	Connections utilits.ExpectedConnections
}

func NewBaseServer(config *internal.Config, connections utilits.ExpectedConnections, logger *log.Logger) *BaseServer {
	return &BaseServer{
		Config:      config,
		Logger:      logger,
		Connections: connections,
	}
}

func (s *BaseServer) Check() error {
	if s.Config == nil || s.Logger == nil {
		return ArgError
	}

	if err := s.checkConnection(); err != nil {
		return err
	}
	return nil
}

func (s *BaseServer) checkConnection() error {
	if err := s.Connections.SqlConnection.Ping(); err != nil {
		return fmt.Errorf("Can't check connection to sql with error %v ", err)
	}

	s.Logger.Info("Success check connection to sql db")

	return nil
}
