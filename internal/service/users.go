package service

import (
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
