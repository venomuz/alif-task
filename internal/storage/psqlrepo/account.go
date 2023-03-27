package psqlrepo

import (
	"context"
	"github.com/google/uuid"
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
		"id",
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

func (a *AccountsRepo) UpdateLastVisit(ctx context.Context, account *models.AccountOut) error {
	columns := map[string]interface{}{
		"last_visit": account.LastVisit,
	}

	err := a.db.WithContext(ctx).Model(models.Accounts{}).Where("id = ?", account.ID).Updates(columns).Error

	return err
}

func (a *AccountsRepo) GetByID(ctx context.Context, ID uuid.UUID) (models.AccountOut, error) {
	var account models.AccountOut

	err := a.db.WithContext(ctx).Model(models.Accounts{}).First(&account, "id = ?", ID).Error

	return account, err
}

func (a *AccountsRepo) GetByPhoneNumber(ctx context.Context, phone string) (models.AccountOut, error) {
	var account models.AccountOut

	err := a.db.WithContext(ctx).Model(models.Accounts{}).First(&account, "phone_number = ?", phone).Error

	return account, err
}

func (a *AccountsRepo) GetAll(ctx context.Context) ([]models.AccountOut, error) {
	//TODO implement me
	panic("implement me")
}
