package account

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AccountController struct {
	AccountService IAccountService
}

func NewAccountController(db *sql.DB) *AccountController {
	accountService := NewAccountService(db)
	return &AccountController{
		AccountService: accountService,
	}
}

// @summary     Create Account
// @description Create an account
// @tags        Account Controller
// @accept		json
// @param 		request body CreateAccountRequest true "request body"
// @success     201 {object} Nil
// @failure		400 {object} errors.ErrorResponse "INVALID_SYS_PARAM"
// @failure		500 {object} errors.ErrorResponse "SYS_INTERNAL_SERVER_ERROR"
// @router      /account [post]
func (ac *AccountController) CreateAccount(c *gin.Context) {
	var body CreateAccountRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		logrus.Error("[BindingError]-", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "Missing Params")
		return
	}

	// Validate Request
	err = ac.AccountService.ValidateCreateAccountReq(body)
	if err != nil {
		logrus.Error("Validation error ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Create new account
	err = ac.AccountService.CreateAccount(body)
	if err != nil {
		logrus.Error("Creation error ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// @summary     Get Account
// @description Get an account by id
// @tags        Account Controller
// @param 		url id
// @success     200 {object} GetAccountResponse
// @failure		400 {object} errors.ErrorResponse "INVALID_SYS_PARAM"
// @failure		500 {object} errors.ErrorResponse "SYS_INTERNAL_SERVER_ERROR"
// @router      /account/:id [get]
func (ac *AccountController) GetAccountByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error("[id error]-", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "input id must be integer")
		return
	}

	record, err := ac.AccountService.GetAccountByID(id)
	if err != nil {
		logrus.Error("[query error]-", err)
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusBadRequest, "account id not found")
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to query")
		return
	}

	c.JSON(http.StatusOK, GetAccountResponse{
		AccountId: record.AccID,
		Balance:   strconv.FormatFloat(float64(record.Balance), 'f', -1, 32),
	})
}
