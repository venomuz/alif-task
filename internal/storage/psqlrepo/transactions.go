package psqlrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/models"
	"gorm.io/gorm"
)

func NewTransactionsRepo(db *gorm.DB) *TransactionsRepo {
	return &TransactionsRepo{
		db: db,
	}
}

type TransactionsRepo struct {
	db *gorm.DB
}

func (t *TransactionsRepo) TopUp(ctx context.Context, input *models.TransactionOut) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {

		wallet := models.Wallets{}
		if err := tx.WithContext(ctx).First(&wallet, "account_id = ?", input.AccountID).Error; err != nil {
			return err
		}

		input.AccountBalance = wallet.Balance

		wallet.Balance += input.Amount

		columns := map[string]interface{}{
			"balance":    wallet.Balance,
			"updated_at": input.CreatedAt,
		}

		if err := tx.WithContext(ctx).Model(wallet).Updates(columns).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(models.Transactions{}).Select("*").Create(input).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionsRepo) TransferByPhoneNumber(ctx context.Context, input *models.TransactionOut) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {

		walletSender := models.Wallets{}
		// Get Sender Wallet
		if err := tx.WithContext(ctx).First(&walletSender, "account_id = ?", input.AccountID).Error; err != nil {
			return err
		}

		input.AccountBalance = walletSender.Balance

		if walletSender.Balance-input.Amount < 0 {
			return models.ErrInsufficient
		}

		columns := map[string]interface{}{
			"balance":    walletSender.Balance - input.Amount,
			"updated_at": input.CreatedAt,
		}

		// Update balance Sender
		if err := tx.WithContext(ctx).Model(walletSender).Updates(columns).Error; err != nil {
			return err
		}

		// Create Sender Transaction
		if err := tx.WithContext(ctx).Model(models.Transactions{}).Select("*").Create(input).Error; err != nil {
			return err
		}

		walletReceiver := models.Wallets{}

		// Get Receiver Wallet
		if err := tx.WithContext(ctx).First(&walletReceiver, "account_id = ?", input.Receiver).Error; err != nil {
			return err
		}

		// Open struct to Create Transaction for Receiver
		transaction := models.TransactionOut{
			ID:             uuid.New(),
			AccountID:      *input.Receiver,
			Method:         "IN",
			Reason:         "RECEIVE FUNDS",
			Amount:         input.Amount,
			AccountBalance: walletReceiver.Balance,
			Sender:         &input.AccountID,
			Receiver:       nil,
			CreatedAt:      input.CreatedAt,
		}

		columns["balance"] = walletReceiver.Balance + input.Amount

		// Update balance Receiver
		if err := tx.WithContext(ctx).Model(walletReceiver).Updates(columns).Error; err != nil {
			return err
		}

		// Create Receiver Transaction
		if err := tx.WithContext(ctx).Model(models.Transactions{}).Select("*").Create(&transaction).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
