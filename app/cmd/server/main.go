package main

import (
	"flag"
	"os"

	utilits_redis "getme-backend/internal/pkg/utilits/redis"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"

	"getme-backend/internal"
	"getme-backend/internal/app/server"
	"getme-backend/internal/pkg/utilits"
)

var (
	configPath string
	sqlDB      string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/server.toml", "path to config file")
	flag.StringVar(&sqlDB, "sql-db", "postgres", "what sql-db the application uses")

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

	db, closeResource, err := utilits.GetSQLConnection(config.Repository, sqlDB, config.DebugMode)
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

	cacheURL := config.Repository.CacheURL
	if config.DebugMode {
		cacheURL = config.Repository.CacheURLLocal
	}

	cacheRedisPool := utilits_redis.NewRedisPool(cacheURL)
	logger.Info("cache redis pool create")

	conn, err := cacheRedisPool.Dial()
	if err != nil {
		logger.Fatal(err)
	}

	if err = conn.Close(); err != nil {
		logger.Fatal(err)
	}

	logger.Info("cache new redis pool success check")

	sessionURL := config.Microservices.SessionServerUrl
	if config.DebugMode {
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
			CacheRedisPool:        cacheRedisPool,
		},
		logger,
	)
	if err = serv.Start(&config); err != nil {
		logger.Fatal(err)
	}

	logger.Info("Server was stopped")
}
