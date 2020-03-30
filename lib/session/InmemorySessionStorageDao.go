package session

import (
	"regexp/syntax"
	"time"
)

type InmemorySessionStorageDao struct {
	Data map[string]string
}

func (r InmemorySessionStorageDao) Store(key, value string, expiration time.Duration) error {
	r.Data[key] = value
	return nil
}

func (r InmemorySessionStorageDao) Find(key string) (*string, error) {
	result, exist := r.Data[key]
	if !exist {
		return nil, &syntax.Error{Code: syntax.ErrInternalError, Expr: ""}
	}
	return &result, nil
}

func (r InmemorySessionStorageDao) Remove(key string) error {
	delete(r.Data, key)
	return nil
}
