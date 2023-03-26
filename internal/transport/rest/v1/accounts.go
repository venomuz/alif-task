package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/pkg/logger"
	"net/http"
	"strconv"
)

func (h *Handler) initAccountsRoutes(v1 *gin.RouterGroup) {
	accounts := v1.Group("accounts")
	{
		accounts.POST("/sing-up", h.AccountSingUp)
		accounts.POST("/sing-in", h.AccountSignIn)
		accounts.PUT("/:id", h.AccountUpdate)
		accounts.GET("", h.AccountsGet)
		accounts.GET("accounts/:id", h.AccountGetByID)
		accounts.DELETE("accounts/:id", h.AccountDelete)
	}
}

// AccountSingUp
//	@Summary		Sing Up an account.
//	@Description	This API to Sing Up an account.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.SignUpAccountInput	true	"data body"
//	@Success		201		{object}	models.Accounts
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/accounts/sing-up [POST]
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
//	@Param			data	body		models.SignUpAccountInput	true	"data body"
//	@Success		201		{object}	models.Accounts
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/accounts/sing-up [POST]
func (h *Handler) AccountSignIn(c *gin.Context) {
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

// AccountUpdate
//	@Summary		Update a account.
//	@Description	This API to update a account.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"id for update Account"	Format(id)
//	@Param			data	body		models.UpdateAccountInput	true	"data body"
//	@Success		200		{object}	models.Accounts
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/accounts/{id} [PUT]
func (h *Handler) AccountUpdate(c *gin.Context) {
	var body models.UpdateAccountInput

	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		logger.Zap.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	err = c.ShouldBind(&body)
	if err != nil || inputID == "" {
		newResponse(c, http.StatusBadRequest, models.ErrInputBody.Error())
		logger.Zap.Error("error while bind to json AccountUpdate", logger.Error(err))
		return
	}

	body.ID = uint32(ID)

	account, err := h.services.Accounts.Update(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while update account", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, &account)
}

// AccountsGet
//	@Summary		Get all Accounts
//	@Description	This api for get Accounts
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.Accounts
//	@Failure		500	{object}	Response
//	@Router			/v1/accounts [GET]
func (h *Handler) AccountsGet(c *gin.Context) {

	accounts, err := h.services.Accounts.GetAll(c.Request.Context())
	if err != nil {
		newResponse(c, http.StatusInternalServerError, models.ErrGetAll.Error())
		logger.Zap.Error("error while get all AccountsGet", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// AccountGetByID
//	@Summary		Get account by id.
//	@Description	This API to get account by id.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"id for get Account"	Format(id)
//	@Success		200		{object}	models.Accounts
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/accounts/{id} [GET]
func (h *Handler) AccountGetByID(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		logger.Zap.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	account, err := h.services.Accounts.GetByID(c.Request.Context(), uint32(ID))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, account)
}

// AccountGetByUrl
//	@Summary		Get account by url.
//	@Description	This API to get account by url.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			url	path		string	true	"url for get Account"	Format(id)
//	@Success		200	{object}	models.Accounts
//	@Failure		400	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/v1/accounts-url/{url} [GET]
func (h *Handler) AccountGetByUrl(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		logger.Zap.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	account, err := h.services.Accounts.GetByID(c.Request.Context(), uint32(ID))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, account)
}

// AccountDelete this api deletes Account
//	@Summary		Delete a account.
//	@Description	This API to delete a account.
//	@Tags			Accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id for delete Account"	Format(id)
//	@Success		200	{object}	Response
//	@Router			/v1/accounts/{id} [DELETE]
func (h *Handler) AccountDelete(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		logger.Zap.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	err = h.services.Accounts.DeleteByID(c.Request.Context(), uint32(ID))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "success"})
}
