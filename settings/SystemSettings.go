package settings

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() (string, error) {
	uuidObj, err := uuid.NewUUID()
	data := []byte("wnw8olzvmjp0x6j7ur8vafs4jltjabi0")
	uuidObj2 := uuid.NewMD5(uuidObj, data)
	return uuidObj2.String(), err
}

func PasswordHash(plainPass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func VerifyPassword(hash, s string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
}
