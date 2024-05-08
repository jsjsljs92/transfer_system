package transaction

import (
	"database/sql"

	"github.com/jsjsljs92/transferSystem/src/models"
)

type ITransactionRepository interface {
	CreatePayin(payin models.Payin) error
	CreatePayout(payin models.Payout) error
}

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (tr *TransactionRepository) CreatePayin(payin models.Payin) error {
	_, err := tr.DB.Exec("INSERT INTO payin (to_acc_id, amount, timestamp) VALUES ($1, $2, $3);",
		payin.ToAccID,
		payin.Amount,
		payin.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TransactionRepository) CreatePayout(payout models.Payout) error {
	_, err := tr.DB.Exec("INSERT INTO payout (from_acc_id, amount, timestamp) VALUES ($1, $2, $3);",
		payout.FromAccID,
		payout.Amount,
		payout.Timestamp)
	if err != nil {
		return err
	}
	return nil
}
