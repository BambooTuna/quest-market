package session

type SessionStorageDao interface {
	Store(key, value string) error
	Find(key string) (*string, error)
	Remove(key string) error
}
