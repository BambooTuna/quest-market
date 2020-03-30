package session

import (
	"errors"
	"time"
)

type InmemorySessionStorageDao struct {
	Data map[string]string
}

func (r InmemorySessionStorageDao) Store(key, value string, expiration time.Duration) error {
	println("Store: "+key, "value: "+value)
	r.Data[key] = value
	return nil
}

func (r InmemorySessionStorageDao) Find(key string) (*string, error) {
	println("Find: " + key)
	result, exist := r.Data[key]
	if !exist {
		return nil, errors.New("forbidden")
	}
	return &result, nil
}

func (r InmemorySessionStorageDao) Remove(key string) error {
	println("Remove: " + key)
	delete(r.Data, key)
	return nil
}
