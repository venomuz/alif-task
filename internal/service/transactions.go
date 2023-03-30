package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/repository/psqlrepo"
	"github.com/venomuz/alif-task/internal/repository/redisrepo"
	"github.com/venomuz/alif-task/pkg/logger"
	"regexp"
	"time"
)

func NewTransactionsService(transactionsRepo psqlrepo.Transactions, accountsRepo psqlrepo.Accounts, rdb redisrepo.Repository) *TransactionsService {
	return &TransactionsService{
		transactionsRepo: transactionsRepo,
		accountsRepo:     accountsRepo,
		rdb:              rdb,
	}
}

type TransactionsService struct {
	transactionsRepo psqlrepo.Transactions
	accountsRepo     psqlrepo.Accounts
	rdb              redisrepo.Repository
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
	if err != nil {
		return models.TransactionOut{}, err
	}

	err = t.rdb.Del(ctx, transaction.AccountID.String()+":wallet")
	if err != nil {
		logger.Zap.Error("error delete wallet from Redis", logger.Error(err))
	}

	return transaction, err
}

func (t *TransactionsService) TransferByPhoneNumber(ctx context.Context, input models.TransferByPhoneNumberInput) (models.TransactionOut, error) {
	if input.PinCode != input.AccountPinCode {
		return models.TransactionOut{}, models.ErrPinCodeWrong
	}

	match, _ := regexp.MatchString("998[0-9]{2}[0-9]{7}$", input.ReceiverPhone)
	if !match {
		return models.TransactionOut{}, models.ErrPhoneNumber
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
	if err != nil {
		return models.TransactionOut{}, err
	}

	err = t.rdb.Del(ctx, transaction.AccountID.String()+":wallet")
	if err != nil {
		logger.Zap.Error("error while delete wallet sender from Redis TransferByPhoneNumber", logger.Error(err))
	}

	err = t.rdb.Del(ctx, transaction.Receiver.String()+":wallet")
	if err != nil {
		logger.Zap.Error("error while delete wallet receiver from Redis TransferByPhoneNumber", logger.Error(err))
	}

	return transaction, err
}

func (t *TransactionsService) WithdrawalFunds(ctx context.Context, input models.WithdrawalFundsInput) (models.TransactionOut, error) {
	if input.PinCode != input.AccountPinCode {
		return models.TransactionOut{}, models.ErrPinCodeWrong
	}

	transaction := models.TransactionOut{
		ID:        uuid.New(),
		AccountID: input.AccountID,
		Method:    "OUT",
		Reason:    "WITHDRAWAL FUNDS",
		Amount:    input.Amount,
		CreatedAt: time.Now(),
	}

	err := t.transactionsRepo.WithdrawalFunds(ctx, &transaction)
	if err != nil {
		return models.TransactionOut{}, err
	}

	err = t.rdb.Del(ctx, transaction.AccountID.String()+":wallet")
	if err != nil {
		logger.Zap.Error("error delete wallet from Redis", logger.Error(err))
	}

	return transaction, err
}
