package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/config"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/storage/psqlrepo"
	"github.com/venomuz/alif-task/pkg/auth"
	"github.com/venomuz/alif-task/pkg/hash"
)

type Accounts interface {
	SingUp(ctx context.Context, input models.SignUpAccountInput) (models.AccountOut, error)
	SingIn(ctx context.Context, input models.SingInAccountInput) (models.Tokens, error)
	Update(ctx context.Context, input models.UpdateAccountInput) (models.AccountOut, error)
	GetByID(ctx context.Context, ID uint32) (models.AccountOut, error)
	GetByAccessToken(ctx context.Context, accessToken string) (models.AccountOut, error)
	GetAll(ctx context.Context) ([]models.AccountOut, error)
}

type Settings interface {
	Create(ctx context.Context, input models.CreateSettingInput) (models.SettingOut, error)
	Update(ctx context.Context, input models.UpdateSettingInput) (models.SettingOut, error)
	GetByID(ctx context.Context, ID uint32) (models.SettingOut, error)
	GetAll(ctx context.Context) ([]models.SettingOut, error)
	GetByKey(ctx context.Context, key string) (models.SettingOut, error)
	DeleteByID(ctx context.Context, ID uint32) error
}

type Transactions interface {
	TopUp(ctx context.Context, input models.TopUpInput) (models.TransactionOut, error)
	TransferByPhoneNumber(ctx context.Context, input models.TransferByPhoneNumberInput) (models.TransactionOut, error)
}

type Wallets interface {
	GetByAccountID(ctx context.Context, accountID uuid.UUID) (models.WalletOut, error)
}

type Deps struct {
	PsqlRepo     *psqlrepo.Repositories
	Cfg          config.Config
	Hash         hash.PasswordHasher
	TokenManager auth.TokenManager
}

type Services struct {
	Accounts     Accounts
	Settings     Settings
	Transactions Transactions
	Wallets      Wallets
}

func NewServices(deps Deps) *Services {
	accountsService := NewAccountsService(deps.PsqlRepo.Accounts, deps.PsqlRepo.Wallets, deps.Hash, deps.TokenManager)
	settingsService := NewSettingsService(deps.PsqlRepo.Settings)
	transactionsService := NewTransactionsService(deps.PsqlRepo.Transactions, deps.PsqlRepo.Accounts)
	walletsService := NewWalletsService(deps.PsqlRepo.Wallets)
	return &Services{
		Accounts:     accountsService,
		Settings:     settingsService,
		Transactions: transactionsService,
		Wallets:      walletsService,
	}
}
