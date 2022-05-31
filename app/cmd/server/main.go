package main

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	"getme-backend/internal"
	"getme-backend/internal/app/server"
	"getme-backend/internal/pkg/utilits"
)

var (
	configPath string
	sqlDB      string
	debugMode  bool
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/server.toml", "path to config file")
	flag.StringVar(&sqlDB, "sql-db", "postgres", "what sql-db the application uses")
	flag.BoolVar(&debugMode, "debug", false, "run in debug mode (local configuration)")

}

func main() {
	flag.Parse()
	logrus.Info(os.Args[:])
	//sqlDB = "mysql"
	config := internal.Config{}

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		logrus.Fatal(err)
	}

	logger, closeResource := utilits.NewLogger(&config, false, "")

	defer func(closer func() error, log *logrus.Logger) {
		err := closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResource, logger)

	db, closeResource, err := utilits.GetSQLConnection(config.Repository, sqlDB, debugMode)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Infof("success connect to %v\n", sqlDB)

	defer func(closer func() error, log *logrus.Logger) {
		err := closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResource, logger)

	sessionURL := config.Microservices.SessionServerUrl
	if debugMode {
		sessionURL = config.Microservices.SessionServerUrlLocal
	}
	sessionConn, err := utilits.NewGrpcConnection(sessionURL)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("success connect to session service")

	serv := server.New(&config,
		utilits.ExpectedConnections{
			SqlConnection:         db,
			SessionGrpcConnection: sessionConn,
		},
		logger,
	)
	if err = serv.Start(&config); err != nil {
		logger.Fatal(err)
	}

	logger.Info("Server was stopped")
}
