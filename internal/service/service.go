package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/config"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/repository/psqlrepo"
	"github.com/venomuz/alif-task/internal/repository/redisrepo"
	"github.com/venomuz/alif-task/pkg/auth"
	"github.com/venomuz/alif-task/pkg/hash"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Accounts interface {
	SingUp(ctx context.Context, input models.SignUpAccountInput) (models.AccountOut, error)
	SingIn(ctx context.Context, input models.SingInAccountInput) (models.Tokens, error)
	Update(ctx context.Context, input models.UpdateAccountInput) (models.AccountOut, error)
	GetByID(ctx context.Context, ID uint32) (models.AccountOut, error)
	GetByAccessToken(ctx context.Context, accessToken string) (models.AccountOut, error)
}

type Transactions interface {
	TopUp(ctx context.Context, input models.TopUpInput) (models.TransactionOut, error)
	TransferByPhoneNumber(ctx context.Context, input models.TransferByPhoneNumberInput) (models.TransactionOut, error)
	WithdrawalFunds(ctx context.Context, input models.WithdrawalFundsInput) (models.TransactionOut, error)
}

type Wallets interface {
	GetByAccountID(ctx context.Context, accountID uuid.UUID) (models.WalletOut, error)
}

type Deps struct {
	PsqlRepo     *psqlrepo.Repositories
	RedisRepo    redisrepo.Repository
	Cfg          config.Config
	Hash         hash.PasswordHasher
	TokenManager auth.TokenManager
}

type Services struct {
	Accounts     Accounts
	Transactions Transactions
	Wallets      Wallets
}

func NewServices(deps Deps) *Services {
	accountsService := NewAccountsService(deps.PsqlRepo.Accounts, deps.PsqlRepo.Wallets, deps.RedisRepo, deps.Hash, deps.TokenManager)
	transactionsService := NewTransactionsService(deps.PsqlRepo.Transactions, deps.PsqlRepo.Accounts, deps.RedisRepo)
	walletsService := NewWalletsService(deps.PsqlRepo.Wallets, deps.RedisRepo)
	return &Services{
		Accounts:     accountsService,
		Transactions: transactionsService,
		Wallets:      walletsService,
	}
}
