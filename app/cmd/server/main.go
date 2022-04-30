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
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()
	logrus.Info(os.Args[:])

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

	db, closeResource := utilits.NewPostgresConnection(config.Repository.DataBaseUrl)

	defer func(closer func() error, log *logrus.Logger) {
		err := closer()
		if err != nil {
			log.Fatal(err)
		}
	}(closeResource, logger)

	serv := server.New(&config,
		utilits.ExpectedConnections{
			SqlConnection: db,
		},
		logger,
	)
	if err = serv.Start(&config); err != nil {
		logger.Fatal(err)
	}
	logger.Info("Server was stopped")
}
