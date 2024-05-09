package account

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/jsjsljs92/transferSystem/src/models"
)

type IAccountService interface {
	ValidateCreateAccountReq(req CreateAccountRequest) error
	CreateAccount(req CreateAccountRequest) error
	GetAccountByID(id int) (*models.Account, error)
}

type AccountService struct {
	DB                *sql.DB
	AccountRepository IAccountRepository
}

func NewAccountService(db *sql.DB) *AccountService {
	accountRepository := NewAccountRepository(db)
	return &AccountService{
		DB:                db,
		AccountRepository: accountRepository,
	}
}

func (as *AccountService) ValidateCreateAccountReq(req CreateAccountRequest) error {

	balance, err := convertToFloat32(req.Balance)
	if err != nil {
		return errors.New("balance must be make up of numbers")
	}

	// if balance is negative
	if balance < 0 {
		return errors.New("initial balance cannot be negative")
	}

	// check if account id already exist
	record, err := as.AccountRepository.GetAccountByID(req.AccountId)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}
	if record != nil && record.AccID == req.AccountId {
		return errors.New("account id already exist")
	}

	return nil
}

func (as *AccountService) CreateAccount(req CreateAccountRequest) error {

	err := as.AccountRepository.CreateAccount(req)
	if err != nil {
		return err
	}
	return nil
}

func (as *AccountService) GetAccountByID(id int) (*models.Account, error) {
	// Retrieve an account by ID
	record, err := as.AccountRepository.GetAccountByID(id)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func convertToFloat32(value string) (float32, error) {
	// Convert string to float64
	floatVal, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}

	// Format float to have at most 5 decimal points
	formattedVal := float32(floatVal)
	return formattedVal, nil
}
