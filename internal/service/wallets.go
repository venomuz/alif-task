package service

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/repository/psqlrepo"
	"github.com/venomuz/alif-task/internal/repository/rdb"
	"time"
)

func NewWalletsService(walletsRepo psqlrepo.Wallets, rdb rdb.Repository) *WalletsService {
	return &WalletsService{
		walletsRepo: walletsRepo,
		rdb:         rdb,
	}
}

type WalletsService struct {
	walletsRepo psqlrepo.Wallets
	rdb         rdb.Repository
}

func (w *WalletsService) GetByAccountID(ctx context.Context, accountID uuid.UUID) (models.WalletOut, error) {
	var wallet models.WalletOut

	val, err := w.rdb.Get(ctx, accountID.String()+":wallet")
	if err == nil {
		err = json.Unmarshal([]byte(val), &wallet)
		if err != nil {
			return models.WalletOut{}, err
		}

		return wallet, nil
	}

	wallet, err = w.walletsRepo.GetByAccountID(ctx, accountID)

	walletJson, err := json.Marshal(wallet)
	if err != nil {
		return models.WalletOut{}, err
	}

	err = w.rdb.SetEX(ctx, accountID.String()+":wallet", walletJson, time.Minute*3)
	if err != nil {
		return models.WalletOut{}, err
	}

	return wallet, nil
}
