package settings

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
)

func GenerateUUID() (string, error) {
	uuidObj, err := uuid.NewUUID()
	data := []byte("wnw8olzvmjp0x6j7ur8vafs4jltjabi0")
	uuidObj2 := uuid.NewMD5(uuidObj, data)
	return uuidObj2.String(), err
}

func PasswordHash(plainPass string) (string, error) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(plainPass))), nil
}

func VerifyPassword(hash, s string) error {
	if hash == fmt.Sprintf("%x", sha256.Sum256([]byte(s))) {
		return nil
	} else {
		return errors.New("forbidden")
	}
}

func FetchEnvValue(key string, defaultValue string) string {
	dataSourceName := os.Getenv(key)
	if dataSourceName == "" {
		dataSourceName = defaultValue
	}
	return dataSourceName
}
