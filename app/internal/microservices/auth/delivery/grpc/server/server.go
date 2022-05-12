package server

import (
	"context"
	"net"

	proto "getme-backend/internal/microservices/auth/delivery/grpc/protobuf"
	"getme-backend/internal/microservices/auth/sessions"
	//prometheus_monitoring "getme-backend/pkg/monitoring/prometheus-monitoring"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type AuthServer struct {
	grpcServer     *grpc.Server
	sessionUsecase sessions.SessionUsecase
	logger         *logrus.Logger
}

func NewAuthGRPCServer(logger *logrus.Logger, grpcServer *grpc.Server, sessionUsecase sessions.SessionUsecase) *AuthServer {
	server := &AuthServer{
		sessionUsecase: sessionUsecase,
		grpcServer:     grpcServer,
		logger:         logger,
	}
	return server
}

func (server *AuthServer) StartGRPCServer(listenUrl string) error {
	lis, err := net.Listen("tcp", listenUrl)
	server.logger.Infof("my listen url %s \n", listenUrl)

	if err != nil {
		server.logger.Errorf("AUTHSERVER\n")
		server.logger.Errorf("can not listen url: %s err :%v\n", listenUrl, err)
		return err
	}
	proto.RegisterAuthCheckerServer(server.grpcServer, server)

	//go prometheus_monitoring.CreateNewMonitoringRouter("sessions-service")

	server.logger.Info("Start session service\n")
	return server.grpcServer.Serve(lis)
}

func (s *AuthServer) Check(_ context.Context, sessionID *proto.SessionID) (*proto.Result, error) {
	s.logger.Infof("AUTHSERVER - Check: call with sessionID = %v\n", sessionID.ID)
	res, err := s.sessionUsecase.Check(sessionID.ID)
	if err != nil {
		s.logger.Errorf("AUTHSERVER\n")
		s.logger.Errorf("can not check session with sessionID = %s, err = %v", sessionID.ID,
			err)
		return nil, err
	}
	s.logger.Infof("AUTHSERVER - Check: correctly work, res = %v\n", res)

	return &proto.Result{
		UserID:    res.UserID,
		SessionID: res.UniqID,
	}, nil
}

func (s *AuthServer) Create(_ context.Context, userID *proto.UserID) (*proto.Result, error) {
	s.logger.Infof("AUTHSERVER - Create: call with userID = %v\n", userID.ID)
	res, err := s.sessionUsecase.Create(userID.ID)
	if err != nil {
		s.logger.Errorf("AUTHSERVER\n")
		s.logger.Errorf("can not create session with userID = %d, err = %v", userID.ID,
			err)
		return nil, err
	}
	s.logger.Infof("AUTHSERVER - Create: correctly work, res = %v\n", res)

	return &proto.Result{
		UserID:    res.UserID,
		SessionID: res.UniqID,
	}, nil
}
func (s *AuthServer) Delete(_ context.Context, sessionID *proto.SessionID) (*proto.Nothing, error) {
	s.logger.Infof("AUTHSERVER - Delete: call with sessionID = %v\n", sessionID.ID)
	err := s.sessionUsecase.Delete(sessionID.ID)
	if err != nil {
		s.logger.Errorf("AUTHSERVER\n")
		s.logger.Errorf("can not delete session with sessionID = %s, err = %v", sessionID.ID,
			err)
		return &proto.Nothing{Dummy: false}, err
	}
	s.logger.Infof("AUTHSERVER - Delete: correctly work\n")

	return &proto.Nothing{
		Dummy: true,
	}, nil
}

func (s *AuthServer) CreateByToken(_ context.Context, data *proto.TokenUserData) (*proto.ResultByToken, error) {
	s.logger.Infof("AUTHSERVER - CreateByToken: call with tokenID = %v and UserID = %v\n", data.Token.ID, data.User.ID)
	res, err := s.sessionUsecase.CreateByTokenID(data.Token.ID, data.User.ID)
	if err != nil {
		s.logger.Errorf("AUTHSERVER\n")
		s.logger.Errorf("can not create session with tokenID = %v and UserID = %v - err = %v", data.Token.ID,
			data.User.ID, err)
		return nil, err
	}
	s.logger.Infof("AUTHSERVER - CreateByToken: correctly work, res = %v\n", res)

	return &proto.ResultByToken{
		Token:  res.TokenID,
		UserID: res.UserID,
	}, nil
}
func (s *AuthServer) CheckWithDelete(_ context.Context, sessionID *proto.TokenID) (*proto.ResultByToken, error) {
	s.logger.Infof("AUTHSERVER - CheckWithDelete: call with sessionID = %v\n", sessionID.ID)
	res, err := s.sessionUsecase.CheckWithDelete(sessionID.ID)
	if err != nil {
		s.logger.Errorf("AUTHSERVER\n")
		s.logger.Errorf("can not check session with sessionID = %v, err = %v", sessionID.ID,
			err)
		return nil, err
	}
	s.logger.Infof("AUTHSERVER - CheckWithDelete: correctly work, res = %v\n", res)

	return &proto.ResultByToken{
		Token:  res.TokenID,
		UserID: res.UserID,
	}, nil
}
