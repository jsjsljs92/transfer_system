package transaction

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jsjsljs92/transferSystem/common"
	account "github.com/jsjsljs92/transferSystem/src/components/account"
	"github.com/jsjsljs92/transferSystem/src/models"
	"github.com/sirupsen/logrus"
)

type ValidateObj struct {
	SourceAcc, DestinationAcc *models.Account
	Amount                    float32
	Err                       error
}

type ITransactionService interface {
	ValidateTransaction(req TransactionRequest) ValidateObj
	CreateTransactionRecord(req TransactionRequest)
	UpdateAccountBalance(accounts ValidateObj) error
}

type TransactionService struct {
	DB                    *sql.DB
	TransactionRepository ITransactionRepository
	AccountService        account.IAccountService
}

func NewTransactionService(db *sql.DB) *TransactionService {
	transactionRepository := NewTransactionRepository(db)
	accountService := account.NewAccountService(db)
	return &TransactionService{
		DB:                    db,
		TransactionRepository: transactionRepository,
		AccountService:        accountService,
	}
}

func (ts *TransactionService) ValidateTransaction(req TransactionRequest) ValidateObj {
	amount, err := common.ConvertToFloat32(req.Amount)
	if err != nil {
		return ValidateObj{
			Err: errors.New("amount must be make up of numbers"),
		}
	}

	// if amount is negative
	if amount < 0 {
		return ValidateObj{
			Err: errors.New("amount cannot be negative"),
		}
	}

	sourceAcc, err := ts.AccountService.GetAccountByID(req.SourceAccountId)
	if err != nil {
		return ValidateObj{
			Err: errors.New("source account id not found"),
		}
	}

	// if source account balance lesser than amount
	if sourceAcc.Balance < amount {
		return ValidateObj{
			Err: errors.New("source account's balance lesser than amount"),
		}
	}

	destinationAcc, err := ts.AccountService.GetAccountByID(req.DestinationAccountId)
	if err != nil {
		return ValidateObj{
			Err: errors.New("destination account id not found"),
		}
	}

	return ValidateObj{
		SourceAcc:      sourceAcc,
		DestinationAcc: destinationAcc,
		Amount:         amount,
	}
}

func (ts *TransactionService) CreateTransactionRecord(req TransactionRequest) {
	amount, _ := common.ConvertToFloat32(req.Amount)
	now := time.Now()

	go func(payin models.Payin) {
		err := ts.TransactionRepository.CreatePayin(payin)
		if err != nil {
			logrus.Error("Failed to create payin: ", err.Error())
		}
	}(models.Payin{
		ToAccID:   req.DestinationAccountId,
		Amount:    amount,
		Timestamp: now,
	})

	go func(payout models.Payout) {
		err := ts.TransactionRepository.CreatePayout(payout)
		if err != nil {
			logrus.Error("Failed to create payout: ", err.Error())
		}
	}(models.Payout{
		FromAccID: req.SourceAccountId,
		Amount:    amount,
		Timestamp: now,
	})
}

func (ts *TransactionService) UpdateAccountBalance(accounts ValidateObj) error {

	err := ts.TransactionRepository.UpdateTransactionBalance(*accounts.SourceAcc, *accounts.DestinationAcc)
	if err != nil {
		return err
	}
	return nil
}
