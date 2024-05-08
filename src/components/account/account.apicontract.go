package account

type CreateAccountRequest struct {
	AccountId int    `json:"account_id" binding:"required"`
	Balance   string `json:"balance" binding:"required"`
}

type GetAccountResponse struct {
	AccountId int    `json:"account_id" binding:"required"`
	Balance   string `json:"balance" binding:"required"`
}
