package transaction

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TransactionController struct {
	TransactionService ITransactionService
}

func NewTransactionController(db *sql.DB) *TransactionController {
	transactionService := NewTransactionService(db)
	return &TransactionController{
		TransactionService: transactionService,
	}
}

// @summary     Create Transaction
// @description Transfer amount from one account to another
// @tags        Transaction Controller
// @accept		json
// @param 		request body TransactionRequest true "request body"
// @success     201 {object} Nil
// @failure		400 {object} errors.ErrorResponse "INVALID_SYS_PARAM"
// @failure		500 {object} errors.ErrorResponse "SYS_INTERNAL_SERVER_ERROR"
// @router      /transactions [post]
func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var body TransactionRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		logrus.Error("[BindingError]-", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "Missing Params")
		return
	}

	// Perform validation
	accounts := tc.TransactionService.ValidateTransaction(body)
	if accounts.Err != nil {
		logrus.Error("Validation error ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, accounts.Err.Error())
		return
	}

	// run transaction creation in another thread
	go func(body TransactionRequest) {
		tc.TransactionService.CreateTransactionRecord(body)
	}(body)

	// amend the values in model
	now := time.Now()
	accounts.SourceAcc.Balance -= accounts.Amount
	accounts.SourceAcc.Timestamp = now
	accounts.DestinationAcc.Balance += accounts.Amount
	accounts.DestinationAcc.Timestamp = now

	err = tc.TransactionService.UpdateAccountBalance(accounts)
	if err != nil {
		logrus.Error("Update error ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
