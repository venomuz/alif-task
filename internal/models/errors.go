package models

import "errors"

var (
	ErrInputBody              = errors.New("invalid input body")
	ErrPhoneOrPasswordWrong   = errors.New("phone number or password wrong")
	ErrPhoneNumber            = errors.New("not valid phone number")
	ErrUnauthorized           = errors.New("you are not logged in")
	ErrNotFoundAccountFromCtx = errors.New("account is missing from ctx")
	ErrPinCodeWrong           = errors.New("you inserted wrong pin code")
	ErrNotFoundAccount        = errors.New("no account with such phone number was found")
	ErrInsufficient           = errors.New("insufficient funds in the wallet")
)

type EmptyStruct struct{}
