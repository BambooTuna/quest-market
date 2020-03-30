package session

import "time"

type RedisSessionStorageDao struct {
}

func (r RedisSessionStorageDao) Store(key, value string, expiration time.Duration) error {
	// TODO
	return nil
}

func (r RedisSessionStorageDao) Find(key string) (*string, error) {
	// TODO
	result := "result value"
	return &result, nil
}

func (r RedisSessionStorageDao) Remove(key string) error {
	// TODO
	return nil
}
