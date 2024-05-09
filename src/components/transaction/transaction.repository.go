package transaction

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jsjsljs92/transferSystem/src/models"
	"github.com/sirupsen/logrus"
)

type ITransactionRepository interface {
	CreatePayin(payin models.Payin) error
	CreatePayout(payin models.Payout) error
	UpdateTransactionBalance(source, destination models.Account) error
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

func (tr *TransactionRepository) UpdateTransactionBalance(source, destination models.Account) error {

	tx, err := tr.DB.Begin()
	if err != nil {
		logrus.Fatal(err)
	}

	query := "UPDATE account SET balance = $1, version = $2, timestamp = $3 WHERE acc_id = $4 AND version = $5"
	_, err = tx.Exec(query, strconv.FormatFloat(float64(source.Balance), 'f', -1, 32),
		source.Version+1,
		source.Timestamp,
		source.AccID,
		source.Version)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update account: %v", err)
	}

	_, err = tx.Exec(query, strconv.FormatFloat(float64(destination.Balance), 'f', -1, 32),
		destination.Version+1,
		destination.Timestamp,
		destination.AccID,
		destination.Version)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update account: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		logrus.Fatal(err)
	}

	return nil
}
