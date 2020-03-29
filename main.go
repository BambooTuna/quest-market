package main

import (
	"fmt"
	"github.com/BambooTuna/quest-market/controller"
	"github.com/BambooTuna/quest-market/dao"
	"github.com/BambooTuna/quest-market/lib/session"
	"github.com/BambooTuna/quest-market/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	sessionDao := session.InmemorySessionStorageDao{Data: make(map[string]string)}
	session := session.DefaultSession{
		Dao: sessionDao,
	}

	accountCredentialsDao := dao.AccountCredentialsDaoImpl{}

	authenticationUseCase := usecase.AuthenticationUseCase{AccountCredentialsDao: accountCredentialsDao}

	authenticationController := controller.AuthenticationController{
		Session:               session,
		AuthenticationUseCase: authenticationUseCase,
	}

	r := gin.Default()
	r.POST("/signup", authenticationController.SignUp())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}
