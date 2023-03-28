package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/pkg/logger"
	"net/http"
)

func (h *Handler) initAccountsRoutes(v1 *gin.RouterGroup) {
	accounts := v1.Group("accounts")

	accounts.POST("/sing-up", h.AccountSingUp)
	accounts.POST("/sing-in", h.AccountSignIn)
	accounts.POST("/refresh", h.AccountRefresh)

	authenticated := accounts.Group("/", h.AccountIdentity)

	{
		authenticated.PUT("", h.AccountUpdate)
		authenticated.GET("/me", h.AccountGetMe)

		wallets := authenticated.Group("/wallets")
		{
			wallets.GET("/balance", h.AccountGetWallet)
			wallets.POST("/top-up", h.AccountWalletTopUp)
			wallets.POST("/transfer-by-phone", h.AccountFundTransfer)
		}
	}
}

// AccountSingUp
//	@Summary		Sing Up an account.
//	@Description	This API to Sing Up an account.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.SignUpAccountInput	true	"data body"
//	@Success		201		{object}	models.AccountOut
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/api/v1/accounts/sing-up [POST]
func (h *Handler) AccountSingUp(c *gin.Context) {
	var body models.SignUpAccountInput

	err := c.ShouldBind(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		logger.Zap.Error("error while bind to json AccountSingUp", logger.Error(err))
		return
	}

	account, err := h.services.Accounts.SingUp(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		logger.Zap.Error("error while create account", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, account)
}

// AccountSignIn
//	@Summary		Sing Up an account.
//	@Description	This API to Sing Up an account.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.SingInAccountInput	true	"data body"
//	@Success		200		{object}	models.Tokens
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/api/v1/accounts/sing-in [POST]
func (h *Handler) AccountSignIn(c *gin.Context) {
	var body models.SingInAccountInput

	err := c.ShouldBind(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		logger.Zap.Error("error while bind to json AccountSingIn", logger.Error(err))
		return
	}

	tokens, err := h.services.Accounts.SingIn(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		logger.Zap.Error("error while create account", logger.Error(err))
		return
	}

	c.SetCookie("access_token", tokens.AccessToken, int(h.cfg.AUTH.AccessTokenTTL.Seconds()), "/", h.cfg.HTTP.Host, false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(h.cfg.AUTH.RefreshTokenTTL.Seconds()), "/", h.cfg.HTTP.Host, false, true)

	c.JSON(http.StatusOK, models.Tokens{AccessToken: tokens.AccessToken})
}

func (h *Handler) AccountRefresh(c *gin.Context) {

}

// AccountUpdate
//	@Summary		Update an account.
//	@Description	This API to update an account.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.UpdateAccountInput	true	"data body"
//	@Success		200		{object}	models.AccountOut
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/api/v1/accounts/ [PUT]
func (h *Handler) AccountUpdate(c *gin.Context) {
	var body models.UpdateAccountInput

	currentAccount, err := h.GetAccountFromCtx(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while get account from ctx", logger.Error(err))
		return
	}

	err = c.ShouldBind(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		logger.Zap.Error("error while bind to json AccountUpdate", logger.Error(err))
		return
	}

	body.ID = currentAccount.ID

	account, err := h.services.Accounts.Update(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while update account", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, account)
}

// AccountGetMe
//	@Summary		Get account by id.
//	@Description	This API to get account by id.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.AccountOut
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/api/v1/accounts/me [GET]
func (h *Handler) AccountGetMe(c *gin.Context) {
	currentAccount, err := h.GetAccountFromCtx(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while get account from ctx", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, currentAccount)
}

// AccountGetWallet
//	@Summary		Get wallet by accountId.
//	@Description	This API to get wallet by accountId.
//	@Tags			Wallets
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.WalletOut
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/api/v1/accounts/wallets/balance [GET]
func (h *Handler) AccountGetWallet(c *gin.Context) {
	currentAccount, err := h.GetAccountFromCtx(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while get account from ctx", logger.Error(err))
		return
	}

	wallet, err := h.services.Wallets.GetByAccountID(c.Request.Context(), currentAccount.ID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while get wallet by account id", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, wallet)
}

// AccountWalletTopUp
//	@Summary		Top up wallet balance.
//	@Description	This API to top up wallet balance.
//	@Tags			Wallets
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.TopUpInput	true	"data body"
//	@Success		200		{object}	models.TransactionOut
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/api/v1/accounts/wallets/top-up [POST]
func (h *Handler) AccountWalletTopUp(c *gin.Context) {
	var body models.TopUpInput

	currentAccount, err := h.GetAccountFromCtx(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while get account from ctx", logger.Error(err))
		return
	}

	err = c.ShouldBind(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		logger.Zap.Error("error while bind to json AccountWalletTopUp", logger.Error(err))
		return
	}

	body.AccountID = currentAccount.ID

	body.AccountPinCode = currentAccount.PinCode

	transaction, err := h.services.Transactions.TopUp(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while update account", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// AccountFundTransfer
//	@Summary		Top up transfer funds to account balance ny phone number.
//	@Description	This API to top up wallet balance.
//	@Tags			Wallets
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.TransferByPhoneNumberInput	true	"data body"
//	@Success		200		{object}	models.TransactionOut
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/api/v1/accounts/wallets/transfer-by-phone [POST]
func (h *Handler) AccountFundTransfer(c *gin.Context) {
	var body models.TransferByPhoneNumberInput

	currentAccount, err := h.GetAccountFromCtx(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while get account from ctx", logger.Error(err))
		return
	}

	err = c.ShouldBind(&body)
	if err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		logger.Zap.Error("error while bind to json AccountWalletTopUp", logger.Error(err))
		return
	}

	body.AccountID = currentAccount.ID

	body.AccountPinCode = currentAccount.PinCode

	transaction, err := h.services.Transactions.TransferByPhoneNumber(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while update account", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, transaction)
}
