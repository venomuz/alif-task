package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/pkg/logger"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"

	accessTokenCookie = "access_token"

	accountCtx = "account"
)

func (h *Handler) AccountIdentity(c *gin.Context) {
	var accessToken string
	cookie, err := c.Cookie(accessTokenCookie)
	fmt.Println(err)
	authHeader := c.Request.Header.Get(authorizationHeader)

	fields := strings.Fields(authHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	} else if err == nil {
		accessToken = cookie
	}

	if accessToken == "" {
		newResponse(c, http.StatusUnauthorized, models.ErrUnauthorized.Error())
		return
	}

	account, err := h.services.Accounts.GetByAccessToken(c.Request.Context(), accessToken)
	if err != nil {
		logger.Zap.Error("error while get account by token", logger.Error(err))
		newResponse(c, http.StatusUnauthorized, models.ErrUnauthorized.Error())
		return
	}

	c.Set(accountCtx, account)
}

func (h *Handler) GetAccountFromCtx(c *gin.Context) (models.AccountOut, error) {
	value, ex := c.Get(accountCtx)
	if !ex {
		return models.AccountOut{}, models.ErrNotFoundAccountFromCtx
	}

	account, ok := value.(models.AccountOut)
	if !ok {
		return models.AccountOut{}, errors.New("failed to convert value from ctx to models.AccountOut")
	}

	return account, nil
}
