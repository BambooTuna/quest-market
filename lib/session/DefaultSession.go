package session

import (
	"github.com/BambooTuna/quest-market/json"
	"github.com/BambooTuna/quest-market/settings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomClaims struct {
	Data string `json:"data"`
	jwt.StandardClaims
}

type DefaultSession struct {
	Dao      SessionStorageDao
	Settings SessionSettings
}

func (s DefaultSession) SetSession(f func(*gin.Context) (*string, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		plainToken, err := f(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		id, err := settings.GenerateUUID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, json.ErrorMessageJson{Message: err.Error()})
			return
		}

		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &CustomClaims{
			Data:           *plainToken,
			StandardClaims: jwt.StandardClaims{Id: id},
		})

		SignedToken, err := token.SignedString([]byte(s.Settings.Secret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, json.ErrorMessageJson{Message: err.Error()})
			return
		}

		if err := s.Dao.Store(id, SignedToken, s.Settings.ExpirationDate); err != nil {
			c.JSON(http.StatusInternalServerError, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		c.Header(s.Settings.SetAuthHeaderName, SignedToken)
	}
}

func (s DefaultSession) RequiredSession(f func(*gin.Context, string)) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.GetHeader(s.Settings.AuthHeaderName)
		token, err := s.Dao.Find(id)
		if err != nil {
			c.JSON(http.StatusForbidden, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		f(c, *token)
	}
}
