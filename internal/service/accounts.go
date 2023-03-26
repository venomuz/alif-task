package service

import (
	"context"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/storage/psqlrepo"
	"regexp"
	"time"
)

func NewAccountsService(accountsRepo psqlrepo.Accounts) *AccountsService {
	return &AccountsService{accountsRepo: accountsRepo}
}

type AccountsService struct {
	accountsRepo psqlrepo.Accounts
}

func (a *AccountsService) SingUp(ctx context.Context, input models.SignUpAccountInput) (models.AccountOut, error) {
	timeNow := time.Now()

	match, _ := regexp.MatchString("998[0-9]{2}[0-9]{7}$", input.PhoneNumber)
	if !match {
		return models.AccountOut{}, models.ErrPhoneNumber
	}

	hashed, err := a.hash.String(input.Password)
	if err != nil {
		return models.AccountOut{}, err
	}

	account := models.AccountOut{
		Name:        input.Name,
		LastName:    input.LastName,
		PhoneNumber: input.PhoneNumber,
		Birthday:    input.Birthday,
		CreatedAt:   &timeNow,
	}

	err := a.accountsRepo.Create(ctx, &account)
}

func (a *AccountsService) SingIn(ctx context.Context, input models.SingInAccountInput) (models.Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsService) Update(ctx context.Context, input models.UpdateAccountInput) (models.AccountOut, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsService) GetByID(ctx context.Context, ID uint32) (models.AccountOut, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsService) GetAll(ctx context.Context) ([]models.AccountOut, error) {
	//TODO implement me
	panic("implement me")
}
