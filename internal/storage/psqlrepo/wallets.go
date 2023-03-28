package psqlrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/models"
	"gorm.io/gorm"
)

func NewWalletsRepo(db *gorm.DB) *WalletsRepo {
	return &WalletsRepo{
		db: db,
	}
}

type WalletsRepo struct {
	db *gorm.DB
}

func (w *WalletsRepo) Create(ctx context.Context, wallet *models.WalletOut) error {
	err := w.db.WithContext(ctx).Debug().Model(models.Wallets{}).Select(
		"id",
		"account_id",
		"created_at",
	).Create(wallet).Error

	return err
}

func (w *WalletsRepo) GetByAccountID(ctx context.Context, accountID uuid.UUID) (models.WalletOut, error) {
	wallet := models.WalletOut{}

	err := w.db.WithContext(ctx).Model(models.Wallets{}).First(&wallet, "account_id = ?", accountID).Error

	return wallet, err
}
