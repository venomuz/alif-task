package models

import "errors"

var (
	ErrUserAlreadyExists        = errors.New("user with such username already exists")
	ErrPageUrlAlreadyExists     = errors.New("page with such url already exists")
	ErrPhoneAlreadyExists       = errors.New("account with such phone number already exists")
	ErrAuthentication           = errors.New("authentication filed")
	ErrInputBody                = errors.New("invalid input body")
	ErrUsernameOrPassword       = errors.New("username or password wrong")
	ErrPhoneOrPasswordWrong     = errors.New("phone number or password wrong")
	ErrNotSendVerification      = errors.New("verification code was not sent")
	ErrSaveSession              = errors.New("failed to save session")
	ErrNotFoundId               = errors.New("not found id")
	ErrPhoneNumber              = errors.New("not valid phone number")
	ErrVerificationCodeWrong    = errors.New("verification code is wrong please try again")
	ErrFileName                 = errors.New("file name not valid")
	ErrWhileSaveFile            = errors.New("error while save file")
	ErrSettingsKeyAlreadyExists = errors.New("settings with such key already exists")
	ErrGetAll                   = errors.New("error get all")
	ErrWhileUpdate              = errors.New("error while update")
	ErrWhileCreate              = errors.New("error while create")
	ErrInvalidToken             = errors.New("invalid token")
	ErrPageNotFound             = errors.New("page not found")
	ErrWhileSyncIiko            = errors.New("error while sync iiko")
	ErrNotFoundFile             = errors.New("not found file")
	ErrNotFoundLatOrLng         = errors.New("not found latitude or longitude")
	ErrNotFoundTerminalId       = errors.New("not found terminal id")
	ErrNotFoundPhoneNumber      = errors.New("not found phone number")
	ErrNotFoundIsSelfService    = errors.New("not found is_self_service")
	ErrNotFoundTime             = errors.New("not found time")
	ErrNotFoundOrderProducts    = errors.New("not found order products")
)

type EmptyStruct struct{}
