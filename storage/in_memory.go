package storage

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

type InMemoryStorageI interface {
	Set(key, value string, exp time.Duration) error
	Get(key string) (string, error)
}

type storageRedis struct {
	client *redis.Client
}

func NewInMemoryStorage(rdb *redis.Client) InMemoryStorageI {
	return &storageRedis{
		client: rdb,
	}
}

func (r *storageRedis) Set(key, value string, exp time.Duration) error {
	err := r.client.Set(context.Background(), key, value, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *storageRedis) Get(key string) (string, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
