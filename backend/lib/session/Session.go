package session

import (
	"github.com/gin-gonic/gin"
)

type Session interface {
	SetSession(f func(*gin.Context) (*string, error)) func(*gin.Context)
	RequiredSession(f func(*gin.Context, *string)) func(*gin.Context)
	OptionalRequiredSession(f func(*gin.Context, *string)) func(*gin.Context)
	InvalidateSession(f func(*gin.Context)) func(*gin.Context)
}
