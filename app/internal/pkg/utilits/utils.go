package utilits

import (
	"fmt"
	"os"
	"time"

	"getme-backend/internal"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

func NewLogger(config *internal.Config, isService bool, serviceName string) (log *logrus.Logger, closeResource func() error) {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logrus.Fatal(err)
	}

	logger := logrus.New()
	currentTime := time.Now().In(time.UTC)
	var servicePath string
	if isService {
		servicePath = serviceName
	}
	formatted := config.LogAddr + fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second()) + "__" + servicePath + ".log"

	//formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
	//	currentTime.Year(), currentTime.Month(), currentTime.Day(),
	//	currentTime.Hour(), currentTime.Minute(), currentTime.Second()) + "__" + servicePath + ".log"

	f, err := os.OpenFile(formatted, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}

	logger.SetOutput(f)

	logger.Writer()
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger, f.Close
}

func GetSQLConnection(repository internal.RepositoryConnections, database string, localMode bool) (db *sqlx.DB, closeResource func() error, err error) {
	var dbURL string
	switch database {
	case "my-sql":
		dbURL = repository.MySQLDatabaseUrl
		return NewMySQLConnection(dbURL)
	default:
		dbURL = repository.DataBaseUrl
		if localMode {
			dbURL = repository.DataBaseUrlLocal
		}
		return NewPostgresConnection(dbURL)
	}
}
func NewPostgresConnection(databaseUrl string) (db *sqlx.DB, closeResource func() error, err error) {
	db, err = sqlx.Open("postgres", databaseUrl)
	if err != nil {
		return nil, nil, err
	}

	return db, db.Close, nil
}
func NewMySQLConnection(databaseUrl string) (db *sqlx.DB, closeResource func() error, err error) {
	db, err = sqlx.Open("mysql", databaseUrl)
	if err != nil {
		return nil, nil, err
	}

	return db, db.Close, nil
}
