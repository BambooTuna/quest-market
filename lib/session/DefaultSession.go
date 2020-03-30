package session

import (
	"github.com/BambooTuna/quest-market/json"
	"github.com/BambooTuna/quest-market/settings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &jwt.StandardClaims{Id: id})

		SignedToken, err := token.SignedString([]byte(s.Settings.Secret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, json.ErrorMessageJson{Message: err.Error()})
			return
		}

		if err := s.Dao.Store(id, *plainToken, s.Settings.ExpirationDate); err != nil {
			c.JSON(http.StatusInternalServerError, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		c.Header(s.Settings.SetAuthHeaderName, SignedToken)
	}
}

func (s DefaultSession) RequiredSession(f func(*gin.Context, *string)) func(*gin.Context) {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(s.Settings.AuthHeaderName)
		if tokenString == "" {
			c.JSON(http.StatusForbidden, json.ErrorMessageJson{Message: "please set session token to header"})
			return
		}

		var claims jwt.StandardClaims
		_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Settings.Secret), nil
		})
		if err != nil {
			c.JSON(http.StatusForbidden, json.ErrorMessageJson{Message: err.Error()})
			return
		}

		token, err := s.Dao.Find(claims.Id)
		if err != nil {
			c.JSON(http.StatusForbidden, json.ErrorMessageJson{Message: err.Error()})
			return
		}
		f(c, token)
	}
}
