package main

import (
	"database/sql"
	"fmt"
	"github.com/BambooTuna/quest-market/backend/controller"
	"github.com/BambooTuna/quest-market/backend/dao"
	"github.com/BambooTuna/quest-market/backend/lib/session"
	"github.com/BambooTuna/quest-market/backend/model/account"
	"github.com/BambooTuna/quest-market/backend/usecase"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/gorp.v1"
	"log"
	"os"
)

func main() {
	apiVersion := "/v1"

	dataSourceName := os.Getenv("MYSQL_SOURCE")
	if dataSourceName == "" {
		dataSourceName = "BambooTuna:pass@tcp(127.0.0.1:3306)/market"
	}
	db, err := sql.Open("mysql", dataSourceName)
	dbSession := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbSession.AddTableWithName(account.AccountCredentials{}, "account_credentials").SetKeys(false, "account_id")
	defer dbSession.Db.Close()
	if err != nil {
		log.Fatal(err)
	}

	sessionDao := session.InmemorySessionStorageDao{Data: make(map[string]string)}
	authSession := session.DefaultSession{Dao: sessionDao, Settings: session.DefaultSessionSettings("1234567890asdfghjkl")}
	accountCredentialsDao := dao.AccountCredentialsDaoImpl{DBSession: dbSession}
	authenticationUseCase := usecase.AuthenticationUseCase{AccountCredentialsDao: accountCredentialsDao}
	authenticationController := controller.AuthenticationController{
		Session:               authSession,
		AuthenticationUseCase: authenticationUseCase,
	}

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./front/dist", false)))
	r.POST(apiVersion+"/signup", authenticationController.SignUp())
	r.POST(apiVersion+"/signin", authenticationController.SignIn())
	r.GET(apiVersion+"/health", authenticationController.Health())
	r.NoRoute(func(c *gin.Context) {
		c.File("./front/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}
