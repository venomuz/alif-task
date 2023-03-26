package psqlrepo

import (
	"context"
	"github.com/venomuz/alif-task/internal/models"
	"gorm.io/gorm"
)

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

type UsersRepo struct {
	db *gorm.DB
}

func (u *UsersRepo) GetByUsername(ctx context.Context, username string) (models.Users, error) {
	var user models.Users

	err := u.db.WithContext(ctx).First(&user, "username = ?", username).Error

	return user, err
}
