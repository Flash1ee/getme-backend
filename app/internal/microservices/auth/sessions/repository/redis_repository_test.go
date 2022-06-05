package repository

import (
	"bytes"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"

	"getme-backend/internal/microservices/auth/sessions/models"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SuiteTestRepository struct {
	suite.Suite
	redisServer     *miniredis.Miniredis
	redisRepository *RedisRepository
	log             *logrus.Logger
	output          string
}

func (s *SuiteTestRepository) SetupSuite() {
	s.log = logrus.New()
	s.log.SetLevel(logrus.FatalLevel)
	s.output = ""
	s.log.SetOutput(bytes.NewBufferString(s.output))

	var err error
	s.redisServer, err = miniredis.Run()
	require.NoError(s.T(), err)

	addr := s.redisServer.Addr()
	redisConn := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}

	s.redisRepository = NewRedisRepository(redisConn, s.log)
}

func (s *SuiteTestRepository) AfterTest(_, _ string) {
	s.SetupSuite()
	s.output = ""
}

func (s *SuiteTestRepository) TearDownSuite() {
	s.redisServer.Close()
}

func TestRedisRepository(t *testing.T) {
	suite.Run(t, new(SuiteTestRepository))
}

func (s *SuiteTestRepository) TestSet() {
	session := &models.Session{
		UserID:     "1",
		UniqID:     "2",
		Expiration: 84000,
	}

	err := s.redisRepository.Set(session)
	require.NoError(s.T(), err)

	value, err := s.redisServer.Get(session.UniqID)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), value, session.UserID)

	s.redisServer.FastForward(time.Second * 100000)

	_, err = s.redisServer.Get(session.UniqID)
	assert.Equal(s.T(), err, miniredis.ErrKeyNotFound)
	assert.Equal(s.T(), s.output, "")

	s.redisServer.SetError("Error")
	err = s.redisRepository.Set(session)
	assert.Error(s.T(), err)
	s.redisServer.Close()
}

func (s *SuiteTestRepository) TestGetUserID() {
	session := &models.Session{
		UserID:     "1",
		UniqID:     "any hash",
		Expiration: 84000,
	}

	err := s.redisRepository.Set(session)
	require.NoError(s.T(), err)

	var userID string
	userID, err = s.redisRepository.GetUserId(session.UniqID, session.Expiration)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), userID, session.UserID)
	assert.Equal(s.T(), s.output, "")

	s.redisServer.SetError("Error")
	_, err = s.redisRepository.GetUserId(session.UniqID, session.Expiration)
	assert.Error(s.T(), err)

	s.redisServer.Close()
}

func (s *SuiteTestRepository) TestDel() {
	session := &models.Session{
		UserID:     "1",
		UniqID:     "any hash",
		Expiration: 84000,
	}

	err := s.redisRepository.Set(session)
	require.NoError(s.T(), err)

	err = s.redisRepository.Del(session)
	require.NoError(s.T(), err)

	_, err = s.redisServer.Get(session.UniqID)
	assert.Equal(s.T(), err, miniredis.ErrKeyNotFound)
	assert.Equal(s.T(), s.output, "")

	s.redisServer.SetError("Error")
	err = s.redisRepository.Del(session)
	assert.Error(s.T(), err)

	s.redisServer.Close()
}
