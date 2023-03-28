package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/storage/psqlrepo"
	"time"
)

func NewTransactionsService(transactionsRepo psqlrepo.Transactions, accountsRepo psqlrepo.Accounts) *TransactionsService {
	return &TransactionsService{
		transactionsRepo: transactionsRepo,
		accountsRepo:     accountsRepo,
	}
}

type TransactionsService struct {
	transactionsRepo psqlrepo.Transactions
	accountsRepo     psqlrepo.Accounts
}

func (t *TransactionsService) TopUp(ctx context.Context, input models.TopUpInput) (models.TransactionOut, error) {
	if input.PinCode != input.AccountPinCode {
		return models.TransactionOut{}, models.ErrPinCodeWrong
	}
	transaction := models.TransactionOut{
		ID:        uuid.New(),
		AccountID: input.AccountID,
		Method:    "IN",
		Reason:    "TOP UP BALANCE",
		Amount:    input.Amount,
		CreatedAt: time.Now(),
	}

	err := t.transactionsRepo.TopUp(ctx, &transaction)

	return transaction, err
}

func (t *TransactionsService) TransferByPhoneNumber(ctx context.Context, input models.TransferByPhoneNumberInput) (models.TransactionOut, error) {
	if input.PinCode != input.AccountPinCode {
		return models.TransactionOut{}, models.ErrPinCodeWrong
	}

	receiver, err := t.accountsRepo.GetByPhoneNumber(ctx, input.ReceiverPhone)
	if err != nil {
		return models.TransactionOut{}, models.ErrNotFoundAccount
	}

	transaction := models.TransactionOut{
		ID:             uuid.New(),
		AccountID:      input.AccountID,
		Method:         "OUT",
		Reason:         "TRANSFER TO ACCOUNT BY PHONE",
		Amount:         input.Amount,
		AccountBalance: 0,
		Sender:         nil,
		Receiver:       &receiver.ID,
		CreatedAt:      time.Now(),
	}

	err = t.transactionsRepo.TransferByPhoneNumber(ctx, &transaction)

	return transaction, err
}
