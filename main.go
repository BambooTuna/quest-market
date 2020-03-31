package main

import (
	"database/sql"
	"fmt"
	"github.com/BambooTuna/quest-market/backend/controller"
	"github.com/BambooTuna/quest-market/backend/dao"
	"github.com/BambooTuna/quest-market/backend/json"
	"github.com/BambooTuna/quest-market/backend/lib/session"
	"github.com/BambooTuna/quest-market/backend/model/account"
	"github.com/BambooTuna/quest-market/backend/model/goods"
	"github.com/BambooTuna/quest-market/backend/settings"
	"github.com/BambooTuna/quest-market/backend/usecase"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"gopkg.in/gorp.v1"
	"log"
	"net/http"
	"os"
)

func main() {
	apiVersion := "/v1"

	mysqlDataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		settings.FetchEnvValue("MYSQL_USER", "BambooTuna"),
		settings.FetchEnvValue("MYSQL_PASS", "pass"),
		settings.FetchEnvValue("MYSQL_HOST", "127.0.0.1"),
		settings.FetchEnvValue("MYSQL_PORT", "3306"),
		settings.FetchEnvValue("MYSQL_DATABASE", "market"),
	)
	db, err := sql.Open("mysql", mysqlDataSourceName)
	dbSession := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbSession.AddTableWithName(account.AccountCredentials{}, "account_credentials").SetKeys(false, "account_id")
	dbSession.AddTableWithName(goods.ProductDetails{}, "product_details").SetKeys(false, "product_id")
	defer dbSession.Db.Close()
	if err != nil {
		log.Fatal(err)
	}

	redisAddr := fmt.Sprintf("%s:%s",
		settings.FetchEnvValue("REDIS_HOST", "127.0.0.1"),
		settings.FetchEnvValue("REDIS_PORT", "6379"),
	)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	//sessionDao := session.InmemorySessionStorageDao{Data: make(map[string]string)}
	sessionDao := session.RedisSessionStorageDao{Client: redisClient}

	authSession := session.DefaultSession{Dao: sessionDao, Settings: session.DefaultSessionSettings(settings.FetchEnvValue("SESSION_SECRET", "1234567890asdfghjkl"))}
	accountCredentialsDao := dao.AccountCredentialsDaoImpl{DBSession: dbSession}
	productDetailsDao := dao.ProductDetailsDaoImpl{DBSession: dbSession}

	authenticationUseCase := usecase.AuthenticationUseCase{AccountCredentialsDao: accountCredentialsDao}
	productDetailsUseCase := usecase.ProductDetailsUseCase{ProductDetailsDao: productDetailsDao}

	authenticationController := controller.AuthenticationController{
		Session:               authSession,
		AuthenticationUseCase: authenticationUseCase,
	}
	productController := controller.ProductController{
		Session:               authSession,
		ProductDetailsUseCase: productDetailsUseCase,
	}

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./front/dist", false)))

	r.POST(apiVersion+"/signup", authenticationController.SignUpRoute())
	r.POST(apiVersion+"/signin", authenticationController.SignInRoute())
	r.GET(apiVersion+"/health", authenticationController.HealthRoute())
	r.DELETE(apiVersion+"/logout", authenticationController.SignOutRoute())

	r.GET(apiVersion+"/products", productController.GetOpenProductsRoute())
	r.GET(apiVersion+"/product/:productId", productController.GetProductDetailsRoute())
	r.GET(apiVersion+"/products/self", productController.GetMyProductListRoute())
	r.POST(apiVersion+"/product", productController.ExhibitionRoute())
	r.PUT(apiVersion+"/product/:productId", productController.UpdateProductDetailsRoute())

	//r.POST(apiVersion+"/oauth2/signin/line", UnimplementedRoute)
	//r.GET(apiVersion+"/oauth2/signin/line", UnimplementedRoute)

	r.NoRoute(func(c *gin.Context) {
		c.File("./front/dist/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
}

func UnimplementedRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, json.ErrorMessageJson{Message: "UnimplementedRoute"})
}
