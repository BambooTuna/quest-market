package main

import (
	"database/sql"
	"fmt"
	"github.com/BambooTuna/quest-market/controller"
	"github.com/BambooTuna/quest-market/dao"
	"github.com/BambooTuna/quest-market/lib/session"
	"github.com/BambooTuna/quest-market/model/account"
	"github.com/BambooTuna/quest-market/usecase"
	"github.com/gin-gonic/gin"
	"gopkg.in/gorp.v1"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("mysql", "BambooTuna:pass@tcp(127.0.0.1:3306)/market")
	dbSession := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbSession.AddTableWithName(account.AccountCredentials{}, "account_credentials").SetKeys(false, "account_id")
	defer log.Fatal(dbSession.Db.Close())
	if err != nil {
		log.Fatal(err)
	}

	sessionDao := session.InmemorySessionStorageDao{Data: make(map[string]string)}
	session := session.DefaultSession{Dao: sessionDao}
	accountCredentialsDao := dao.AccountCredentialsDaoImpl{DBSession: dbSession}
	authenticationUseCase := usecase.AuthenticationUseCase{AccountCredentialsDao: accountCredentialsDao}
	authenticationController := controller.AuthenticationController{
		Session:               session,
		AuthenticationUseCase: authenticationUseCase,
	}

	r := gin.Default()
	r.POST("/signup", authenticationController.SignUp())
	r.POST("/signin", authenticationController.SignIn())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}
