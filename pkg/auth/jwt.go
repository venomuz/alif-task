package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenManager interface {
	GenerateJwtTokens(sub, phone string) (access, refresh string, err error)
	ExtractClaims(inputToken string) (jwt.MapClaims, error)
}

func NewTokenManager(signingKey string) *JwtHandler {
	return &JwtHandler{SignInKey: signingKey}
}

type JwtHandler struct {
	Sub       string
	Iss       uint32
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	Token     string
	SignInKey string
}

//GenerateJwtTokens ...
func (hand *JwtHandler) GenerateJwtTokens(sub, phone string) (access, refresh string, err error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
	)

	accessToken = jwt.New(jwt.SigningMethodHS256)
	claims = accessToken.Claims.(jwt.MapClaims)
	claims["phone"] = phone
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["iat"] = time.Now().Unix()
	access, err = accessToken.SignedString([]byte(hand.SignInKey))
	if err != nil {
		return "", "", err
	}

	refreshToken = jwt.New(jwt.SigningMethodHS256)
	claims = refreshToken.Claims.(jwt.MapClaims)
	claims["phone"] = phone
	claims["sub"] = sub
	refresh, err = refreshToken.SignedString([]byte(hand.SignInKey))
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (hand *JwtHandler) ExtractClaims(inputToken string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	token, err = jwt.Parse(inputToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(hand.SignInKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, fmt.Errorf("error get user claims from token")
	}

	return claims, nil
}
