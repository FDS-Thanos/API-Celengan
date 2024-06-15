package main

import (
	"api_gateway/handler"
	"api_gateway/proto"

	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))

	addrServiceTrxOpt := client.WithAddress(":1809")
	clientServiceTrx := grpc.NewClient()

	srvTransaction := micro.NewService(
		micro.Client(clientServiceTrx),
	)

	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.GET("/balance", handler.NewAccount().BalanceAccount)

	authRoute := r.Group("/auth")
	// authRoute.POST("/login", handler.Login().AuthLogin)
	authRoute.POST("/logindb", handler.Login().AuthLogin)

	transactionRoute := r.Group("/transaction")
	transactionRoute.GET("/get", func(g *gin.Context) {
		clientResponse, err := proto.NewServiceTransactionService("service-transaction", srvTransaction.Client()).Login(context.Background(), &proto.LoginRequest{
			Username: "admin",
			Password: "admin123",
		}, addrServiceTrxOpt)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		g.JSON(http.StatusOK, gin.H{
			"data": clientResponse,
		})
	})
	transactionRoute.POST("/transfer", handler.NewTransaction().Create)
	transactionRoute.GET("/list-bank", new(handler.ListBankImplement).ReceiveBank)
	transactionRoute.GET("/list-account", new(handler.ListAccountImplement).ReceiveAccount)
	transactionRoute.GET("/check-valid-account/:bankid/:accountid", new(handler.ValidAccountImplement).ValidAccount)
	transactionRoute.POST("/transfer-bank", new(handler.TransferAmountImplement).CreateTransaction)
	transactionRoute.GET("/transaction-history", new(handler.TransactionHistoryImplement).ReceiveTransaction)
	r.Run()
}
