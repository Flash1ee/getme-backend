package main

import (
	"flag"
	"os"

	"google.golang.org/grpc"

	"getme-backend/internal"
	sessionServer "getme-backend/internal/microservices/auth/delivery/grpc/server"
	"getme-backend/internal/microservices/auth/sessions/repository"
	"getme-backend/internal/microservices/auth/sessions/usecase"
	"getme-backend/internal/pkg/utilits"
	//prometheus_monitoring "getme-backend/pkg/monitoring/prometheus-monitoring"
	//"getme-backend/pkg/utils"
	utilits_redis "getme-backend/internal/pkg/utilits/redis"
	//grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
	debugMode  bool
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/server.toml", "path to config file")
	flag.BoolVar(&debugMode, "debug", false, "run in debug mode (local configuration)")

}

func main() {
	flag.Parse()
	config := &internal.Config{}

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}
	logger, CloseLogger := utilits.NewLogger(config, true, "session_microservice")

	defer func() {
		if err := CloseLogger(); err != nil {
			logrus.Fatal(err)
		}
	}()

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		os.Exit(1)
	}
	logger.SetLevel(level)

	sessionRedisURL := config.Microservices.SessionRedisUrl
	if debugMode {
		sessionRedisURL = config.Microservices.SessionRedisUrlLocal
	}
	sessionRedisPool := utilits_redis.NewRedisPool(sessionRedisURL)
	logger.Info("Session-service new redis pool create")

	conn, err := sessionRedisPool.Dial()
	if err != nil {
		logger.Fatal(err)
	}
	if err = conn.Close(); err != nil {
		logger.Fatal(err)
	}
	logger.Info("Session-service new redis pool success check")
	//metrics := prometheus_monitoring.NewPrometheusMetrics("auth")
	//if err = metrics.SetupMonitoring(); err != nil {
	//	logger.Fatal(err)
	//}

	//grpcDurationMetrics := utils.AuthInterceptor(metrics)

	grpcServer := grpc.NewServer(
	//grpc2.UnaryInterceptor(grpcDurationMetrics),
	//grpc2.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	//grpc2.UnaryInterceptor(grpcDurationMetrics))

	//grpc_prometheus.Register(grpc)

	sessionURL := config.Microservices.SessionServerUrl
	if debugMode {
		sessionURL = config.Microservices.SessionServerUrlLocal
	}

	sessionRepository := repository.NewRedisRepository(sessionRedisPool, logger)
	logger.Info("Session-service create repository")
	server := sessionServer.NewAuthGRPCServer(logger, grpcServer, usecase.NewSessionUsecase(sessionRepository))
	if err = server.StartGRPCServer(sessionURL); err != nil {
		logger.Fatalln(err)
	}
	logger.Info("Session-service was stopped")

}
