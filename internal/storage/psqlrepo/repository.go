package psqlrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/models"
	"gorm.io/gorm"
)

type Accounts interface {
	Create(ctx context.Context, account *models.AccountOut) error
	Update(ctx context.Context, account *models.AccountOut) error
	UpdateLastVisit(ctx context.Context, account *models.AccountOut) error
	GetByID(ctx context.Context, ID uuid.UUID) (models.AccountOut, error)
	GetByPhoneNumber(ctx context.Context, phone string) (models.AccountOut, error)
	GetAll(ctx context.Context) ([]models.AccountOut, error)
}

type Transactions interface {
	TopUp(ctx context.Context, input *models.TransactionOut) error
	TransferByPhoneNumber(ctx context.Context, input *models.TransactionOut) error
}

type Wallets interface {
	Create(ctx context.Context, wallet *models.WalletOut) error
	GetByAccountID(ctx context.Context, accountID uuid.UUID) (models.WalletOut, error)
}

type Settings interface {
	Create(ctx context.Context, setting *models.SettingOut) error
	Update(ctx context.Context, setting *models.SettingOut) error
	GetByID(ctx context.Context, ID uint32) (models.SettingOut, error)
	GetAll(ctx context.Context) ([]models.SettingOut, error)
	GetByKey(ctx context.Context, key string) (models.SettingOut, error)
	DeleteByID(ctx context.Context, ID uint32) error
}

type Repositories struct {
	Accounts     Accounts
	Settings     Settings
	Transactions Transactions
	Wallets      Wallets
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Accounts:     NewAccountsRepo(db),
		Settings:     NewSettingsRepo(db),
		Transactions: NewTransactionsRepo(db),
		Wallets:      NewWalletsRepo(db),
	}
}
