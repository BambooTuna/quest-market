package session

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type RedisSessionStorageDao struct {
	Client *redis.Client
}

func (r RedisSessionStorageDao) Store(key, value string, expiration time.Duration) error {
	return r.Client.Set(key, value, expiration).Err()
}

func (r RedisSessionStorageDao) Find(key string) (*string, error) {
	result, err := r.Client.Get(key).Result()
	return &result, err
}

func (r RedisSessionStorageDao) Remove(key string) error {
	return r.Client.Del(key).Err()
}
