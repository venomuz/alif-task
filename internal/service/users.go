package service

import (
	"context"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/storage/psqlrepo"
)

func NewUsersService(usersRepo psqlrepo.Users) *UsersService {
	return &UsersService{
		usersRepo: usersRepo,
	}
}

type UsersService struct {
	usersRepo psqlrepo.Users
}

func (u *UsersService) Login(ctx context.Context, input models.LoginUserInput) (models.Users, error) {

	user, err := u.usersRepo.GetByUsername(ctx, input.Username)
	if err != nil {
		return models.Users{}, err
	}

	return user, err
}
