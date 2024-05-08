package account

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jsjsljs92/transferSystem/src/models"
)

type IAccountRepository interface {
	GetAccountByID(id int) (*models.Account, error)
	CreateAccount(req CreateAccountRequest) error
	UpdateAccount(acc models.Account) error
}

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (ar *AccountRepository) GetAccountByID(id int) (*models.Account, error) {
	var account models.Account
	row := ar.DB.QueryRow("SELECT id, acc_id, balance, version, timestamp FROM ACCOUNT WHERE acc_id = $1", id)
	err := row.Scan(&account.ID, &account.AccID, &account.Balance, &account.Version, &account.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.New("internal server error")
	}
	return &account, nil
}

func (ar *AccountRepository) CreateAccount(req CreateAccountRequest) error {
	_, err := ar.DB.Exec("INSERT INTO account (acc_id, balance, version, timestamp) VALUES ($1, $2, $3, $4)", req.AccountId, req.Balance, 1, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (ar *AccountRepository) UpdateAccount(acc models.Account) error {
	query := "UPDATE account SET balance = $1, version = $2, timestamp = $3 WHERE acc_id = $4 AND version = $5"
	_, err := ar.DB.Exec(query, strconv.FormatFloat(float64(acc.Balance), 'f', -1, 32),
		acc.Version+1,
		acc.Timestamp,
		acc.AccID,
		acc.Version)
	if err != nil {
		return fmt.Errorf("failed to update account: %v", err)
	}
	return nil
}
