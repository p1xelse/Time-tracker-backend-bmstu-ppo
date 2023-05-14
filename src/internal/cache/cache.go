package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	ttl = time.Duration(time.Hour)
)

type StorageRedis struct {
	db *redis.Client

	ttl time.Duration
	ctx context.Context
}

func NewStorageRedis(m *redis.Client) *StorageRedis {
	storage := StorageRedis{
		db:  m,
		ttl: ttl,
	}

	return &storage
}
func (sr *StorageRedis) Set(key string, data interface{}) error {
	rawValue, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("fail to marshal data for cache: %w", err)
	}

	err = sr.db.Set(sr.ctx, key, rawValue, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func (sr *StorageRedis) Get(key string) ([]byte, error) {
	resp, err := sr.db.Get(sr.ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to get data from cache: %w", err)
	}

	return resp, nil
}
