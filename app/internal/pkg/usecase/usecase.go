package usecase

import (
	"encoding/hex"

	"getme-backend/internal/app/cache"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/sha3"
)

type BaseUsecase struct {
	cache.StorageRedis
}

func NewBaseUsecase(pool *redis.Pool) *BaseUsecase {
	return &BaseUsecase{
		StorageRedis: *cache.NewStorageRedis(pool),
	}
}

func (u *BaseUsecase) GenToken(initValue string) string {
	value := append([]byte(initValue), uuid.NewV4().Bytes()...)
	hash := sha3.Sum512(value)

	return hex.EncodeToString(hash[:])
}
