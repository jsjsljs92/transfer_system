package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jsjsljs92/transferSystem/src/components/account"
	"github.com/jsjsljs92/transferSystem/src/components/transaction"
)

func CreateRoutes(route *gin.Engine, db *sql.DB) {
	{
		v1 := route.Group("/v1")

		// create account
		{
			accountApi := v1.Group("/accounts")
			accountController := account.NewAccountController(db)

			accountApi.POST("", accountController.CreateAccount)
			accountApi.GET("/:id", accountController.GetAccountByID)
		}

		// create transaction
		{
			transactionApi := v1.Group("/transactions")
			transactionController := transaction.NewTransactionController(db)

			transactionApi.POST("", transactionController.CreateTransaction)
		}
	}
}
