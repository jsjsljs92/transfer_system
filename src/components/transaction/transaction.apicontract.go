package transaction

type TransactionRequest struct {
	SourceAccountId      int    `json:"source_account_id" binding:"required"`
	DestinationAccountId int    `json:"destination_account_id" binding:"required"`
	Amount               string `json:"amount" binding:"required"`
}
