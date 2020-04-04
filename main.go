package main

import (
	"database/sql"
	"fmt"
	"github.com/BambooTuna/quest-market/backend/aggregate"
	"github.com/BambooTuna/quest-market/backend/controller"
	"github.com/BambooTuna/quest-market/backend/dao"
	"github.com/BambooTuna/quest-market/backend/json"
	"github.com/BambooTuna/quest-market/backend/lib/session"
	"github.com/BambooTuna/quest-market/backend/model/account"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
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
	dbSession.AddTableWithName(transaction.MoneyTransaction{}, "money_transaction").SetKeys(true, "transaction_id")
	dbSession.AddTableWithName(dao.ItemDetails{}, "item_details").SetKeys(false, "item_id")
	dbSession.AddTableWithName(dao.ContractDetails{}, "contract_details").SetKeys(false, "item_id")
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
	moneyTransactionDao := dao.MoneyTransactionDaoImpl{DBSession: dbSession}
	itemContractDao := dao.ItemContractDaoImpl{DBSession: dbSession}

	moneyTransactionAggregates := aggregate.MoneyTransactionAggregates{MoneyTransactionDao: moneyTransactionDao, Aggregates: map[string]*aggregate.MoneyTransactionAggregate{}}

	authenticationUseCase := usecase.AuthenticationUseCase{AccountCredentialsDao: accountCredentialsDao}
	moneyManagementUseCase := usecase.MoneyManagementUseCase{ManagementAccountId: settings.FetchEnvValue("ADMIN_ACCOUNT_ID", "f0c28384-3aa4-3f87-9fba-66a0aa62c504"), MoneyTransactionAggregates: &moneyTransactionAggregates}
	itemContractUseCase := usecase.ItemContractUseCase{ItemContractDao: itemContractDao, MoneyManagementUseCase: &moneyManagementUseCase}

	authenticationController := controller.AuthenticationController{
		Session:                authSession,
		AuthenticationUseCase:  authenticationUseCase,
		MoneyManagementUseCase: moneyManagementUseCase,
	}
	moneyManagementController := controller.MoneyManagementController{
		Session:                authSession,
		MoneyManagementUseCase: moneyManagementUseCase,
	}
	itemContractController := controller.ItemContractController{
		Session:             authSession,
		ItemContractUseCase: itemContractUseCase,
	}

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./front/dist", false)))

	r.POST(apiVersion+"/signup", authenticationController.SignUpRoute())
	r.POST(apiVersion+"/signin", authenticationController.SignInRoute())
	r.GET(apiVersion+"/health", authenticationController.HealthRoute())
	r.DELETE(apiVersion+"/logout", authenticationController.SignOutRoute())

	r.GET(apiVersion+"/items", itemContractController.GetPublicItemContractsRoute())
	r.GET(apiVersion+"/items/my", itemContractController.GetMyItemContractsRoute())

	r.POST(apiVersion+"/item", itemContractController.ExhibitionRoute())
	r.GET(apiVersion+"/item/:itemId", itemContractController.GetPublicItemContractRoute("itemId"))
	r.PUT(apiVersion+"/item/:itemId", itemContractController.EditRoute("itemId"))
	r.PUT(apiVersion+"/item/:itemId/purchase", itemContractController.PurchaseItemRoute("itemId"))
	r.PUT(apiVersion+"/item/:itemId/payment", itemContractController.PaymentOfItemPriceRoute("itemId"))
	r.PUT(apiVersion+"/item/:itemId/receipt", itemContractController.ReceiptConfirmationRoute("itemId"))

	r.GET(apiVersion+"/money", moneyManagementController.GetBalanceRoute())
	r.POST(apiVersion+"/money", moneyManagementController.SendMoneyRoute())

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
