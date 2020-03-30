package session

import "time"

type SessionSettings struct {
	Secret            string
	SetAuthHeaderName string
	AuthHeaderName    string
	ExpirationDate    time.Duration
}

func DefaultSessionSettings(secret string) SessionSettings {
	return SessionSettings{
		Secret:            secret,
		SetAuthHeaderName: "Set-Authorization",
		AuthHeaderName:    "Authorization",
		ExpirationDate:    time.Duration(1) * time.Hour,
	}
}
