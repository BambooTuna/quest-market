package session

import "time"

type SessionStorageDao interface {
	Store(key, value string, expiration time.Duration) error
	Find(key string) (*string, error)
	Remove(key string) error
}
