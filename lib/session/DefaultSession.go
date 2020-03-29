package session

import (
	"github.com/BambooTuna/quest-market/json"
	"github.com/BambooTuna/quest-market/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultSession struct {
	Dao SessionStorageDao
}

func (s DefaultSession) SetSession(f func(*gin.Context) (*string, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		token, err := f(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		id, err := settings.GenerateUUID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		if err := s.Dao.Store(id, *token); err != nil {
			c.JSON(http.StatusInternalServerError, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		c.Header("Set-Authorization", id)
	}
}

func (s DefaultSession) RequiredSession(f func(*gin.Context, string)) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.GetHeader("Authorization")
		token, err := s.Dao.Find(id)
		if err != nil {
			c.JSON(http.StatusForbidden, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		f(c, *token)
	}
}
