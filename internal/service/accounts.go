package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/storage/psqlrepo"
	"github.com/venomuz/alif-task/pkg/auth"
	"github.com/venomuz/alif-task/pkg/hash"
	"regexp"
	"time"
)

func NewAccountsService(accountsRepo psqlrepo.Accounts, hasher hash.PasswordHasher, tokenManager auth.TokenManager) *AccountsService {
	return &AccountsService{
		accountsRepo: accountsRepo,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}

type AccountsService struct {
	accountsRepo psqlrepo.Accounts
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
}

func (a *AccountsService) SingUp(ctx context.Context, input models.SignUpAccountInput) (models.AccountOut, error) {
	timeNow := time.Now()

	match, _ := regexp.MatchString("998[0-9]{2}[0-9]{7}$", input.PhoneNumber)
	if !match {
		return models.AccountOut{}, models.ErrPhoneNumber
	}

	password, err := a.hasher.String(input.Password)
	if err != nil {
		return models.AccountOut{}, err
	}

	account := models.AccountOut{
		ID:          uuid.New(),
		Name:        input.Name,
		LastName:    input.LastName,
		PhoneNumber: input.PhoneNumber,
		Birthday:    input.Birthday,
		Password:    password,
		CreatedAt:   &timeNow,
	}

	err = a.accountsRepo.Create(ctx, &account)

	return account, err
}

func (a *AccountsService) SingIn(ctx context.Context, input models.SingInAccountInput) (models.Tokens, error) {
	match, _ := regexp.MatchString("998[0-9]{2}[0-9]{7}$", input.PhoneNumber)
	if !match {
		return models.Tokens{}, models.ErrPhoneNumber
	}

	account, err := a.accountsRepo.GetByPhoneNumber(ctx, input.PhoneNumber)
	if err != nil {
		return models.Tokens{}, models.ErrPhoneOrPasswordWrong
	}

	err = a.hasher.CheckString(account.Password, input.Password)
	if err != nil {
		return models.Tokens{}, models.ErrPhoneOrPasswordWrong
	}

	accessToken, refreshToken, err := a.tokenManager.GenerateJwtTokens(account.ID.String(), account.PhoneNumber)
	if err != nil {
		return models.Tokens{}, err
	}

	nowTime := time.Now()

	account.LastVisit = &nowTime

	err = a.accountsRepo.UpdateLastVisit(ctx, &account)

	return models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *AccountsService) Update(ctx context.Context, input models.UpdateAccountInput) (models.AccountOut, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsService) GetByID(ctx context.Context, ID uint32) (models.AccountOut, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountsService) GetByAccessToken(ctx context.Context, accessToken string) (models.AccountOut, error) {
	claims, err := a.tokenManager.ExtractClaims(accessToken)
	if err != nil {
		return models.AccountOut{}, err
	}

	id, ok := claims["sub"].(string)
	if !ok {
		return models.AccountOut{}, errors.New("error while parse token to string")
	}

	ID, err := uuid.Parse(id)

	account, err := a.accountsRepo.GetByID(ctx, ID)
	if err != nil {
		return models.AccountOut{}, err
	}

	return account, nil
}

func (a *AccountsService) GetAll(ctx context.Context) ([]models.AccountOut, error) {
	//TODO implement me
	panic("implement me")
}
