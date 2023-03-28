package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/storage/psqlrepo"
)

func NewWalletsService(walletsRepo psqlrepo.Wallets) *WalletsService {
	return &WalletsService{
		walletsRepo: walletsRepo,
	}
}

type WalletsService struct {
	walletsRepo psqlrepo.Wallets
}

func (w *WalletsService) GetByAccountID(ctx context.Context, accountID uuid.UUID) (models.WalletOut, error) {
	return w.walletsRepo.GetByAccountID(ctx, accountID)
}
