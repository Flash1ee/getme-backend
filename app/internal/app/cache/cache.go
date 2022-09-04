package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	ttl = time.Hour
)

type StorageRedis struct {
	master *redis.Pool

	ttl time.Duration
}

func NewStorageRedis(m *redis.Pool) *StorageRedis {
	storage := StorageRedis{
		master: m,
		ttl:    ttl,
	}

	return &storage
}
func (sr *StorageRedis) Set(key string, data interface{}) error {
	conn := sr.master.Get()

	defer conn.Close()

	rawValue, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("fail to marshal data for cache: %w", err)
	}
	ttl := int(ttl.Seconds())

	_, err = conn.Do("SETEX", key, ttl, rawValue)
	if err != nil {
		return err
	}
	return nil
}

func (sr *StorageRedis) Get(key string) ([]byte, error) {
	conn := sr.master.Get()

	defer conn.Close()

	resp, err := conn.Do("GET", key)
	if err != nil {
		return nil, fmt.Errorf("failed to get data from cache: %w", err)
	}

	var rawValue []byte
	switch value := resp.(type) {
	case nil: // value is expired
		return nil, nil
	case []byte:
		rawValue = value
	default:
		return nil, fmt.Errorf("wrong value type for data in cache: %v, key = %s", value, key)
	}

	return rawValue, nil
}
