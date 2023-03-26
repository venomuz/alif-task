package psqlrepo

import (
	"context"
	"github.com/venomuz/alif-task/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewAccountsRepo(db *gorm.DB) *AccountsRepo {
	return &AccountsRepo{
		db: db,
	}
}

type AccountsRepo struct {
	db *gorm.DB
}

func (a *AccountsRepo) Create(ctx context.Context, account *models.AccountOut) error {
	err := a.db.WithContext(ctx).Model(models.Accounts{}).Select(
		"name",
		"last_name",
		"phone_number",
		"password",
		"birthday",
		"created_at",
	).Create(account).Error

	return err
}

func (a *AccountsRepo) Update(ctx context.Context, account *models.AccountOut) error {
	columns := map[string]interface{}{
		"name":       account.Name,
		"last_name":  account.LastName,
		"updated_at": account.UpdatedAt,
	}

	err := a.db.WithContext(ctx).Clauses(clause.Returning{}).Model(models.Accounts{}).Updates(columns).Scan(&account).Error

	return err
}

func (a *AccountsRepo) GetByID(ctx context.Context, ID uint32) (models.AccountOut, error) {
	var account models.AccountOut

	err := a.db.WithContext(ctx).Model(models.Accounts{}).First(&account, "id = ?", ID).Error

	return account, err
}

func (a *AccountsRepo) GetAll(ctx context.Context) ([]models.AccountOut, error) {
	//TODO implement me
	panic("implement me")
}
