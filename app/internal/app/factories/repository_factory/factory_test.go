package repository_factory

import (
	"testing"

	"getme-backend/internal/app"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestFactory(t *testing.T) {
	defer func(t *testing.T) {
		err := recover()
		require.Equal(t, err, nil)
	}(t)

	log := &logrus.Logger{}
	t.Helper()
	db, _, err := sqlmock.Newx()
	if err != nil {
		t.Fatal(err)
	}

	factory := NewRepositoryFactory(log, app.ExpectedConnections{SqlConnection: db})
	factory.GetUserRepository()

}
