package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/pkg/logger"
	"net/http"
	"strconv"
)

func (h *Handler) initSettingsRoutes(api *gin.RouterGroup) {
	{
		api.POST("settings", h.SettingCreate)
		api.PUT("settings/:id", h.SettingUpdate)
		api.GET("settings", h.SettingsGet)
		api.GET("settings/:id", h.SettingGetByID)
		api.DELETE("settings/:id", h.SettingDelete)
	}
}

// SettingCreate
//	@Summary		Create a setting.
//	@Description	This API to create a setting.
//	@Tags			Settings
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.CreateSettingInput	true	"data body"
//	@Success		201		{object}	models.Settings
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/settings [POST]
func (h *Handler) SettingCreate(c *gin.Context) {
	var body models.CreateSettingInput

	err := c.ShouldBind(&body)

	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		logger.Zap.Error("error while bind to json SettingCreate", logger.Error(err))
		return
	}

	setting, err := h.services.Settings.Create(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusConflict, err.Error())
		logger.Zap.Error("error while create setting", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, setting)
}

// SettingUpdate
//	@Summary		Update a setting.
//	@Description	This API to update a setting.
//	@Tags			Settings
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"id for update Setting"	Format(id)
//	@Param			data	body		models.UpdateSettingInput	true	"data body"
//	@Success		200		{object}	models.Settings
//	@Failure		400,409	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/settings/{id} [PUT]
func (h *Handler) SettingUpdate(c *gin.Context) {
	var body models.UpdateSettingInput

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
		logger.Zap.Error("error while bind to json SettingUpdate", logger.Error(err))
		return
	}

	body.ID = uint32(ID)

	setting, err := h.services.Settings.Update(c.Request.Context(), body)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		logger.Zap.Error("error while update setting", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, &setting)
}

// SettingsGet
//	@Summary		Get all Settings
//	@Description	This api for get Settings
//	@Tags			Settings
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.Settings
//	@Failure		500	{object}	Response
//	@Router			/v1/settings [GET]
func (h *Handler) SettingsGet(c *gin.Context) {

	settings, err := h.services.Settings.GetAll(c.Request.Context())
	if err != nil {
		newResponse(c, http.StatusInternalServerError, models.ErrGetAll.Error())
		logger.Zap.Error("error while get all SettingsGet", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, settings)
}

// SettingGetByID
//	@Summary		Get setting by id.
//	@Description	This API to get setting by id.
//	@Tags			Settings
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"id for get Setting"	Format(id)
//	@Success		200		{object}	models.Settings
//	@Failure		400,404	{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/v1/settings/{id} [GET]
func (h *Handler) SettingGetByID(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		logger.Zap.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	setting, err := h.services.Settings.GetByID(c.Request.Context(), uint32(ID))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, setting)
}

// SettingGetByUrl
//	@Summary		Get setting by url.
//	@Description	This API to get setting by url.
//	@Tags			Settings
//	@Accept			json
//	@Produce		json
//	@Param			url	path		string	true	"url for get Setting"	Format(id)
//	@Success		200	{object}	models.Settings
//	@Failure		400	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/v1/settings-url/{url} [GET]
func (h *Handler) SettingGetByUrl(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		logger.Zap.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	setting, err := h.services.Settings.GetByID(c.Request.Context(), uint32(ID))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, setting)
}

// SettingDelete this api deletes Setting
//	@Summary		Delete a setting.
//	@Description	This API to delete a setting.
//	@Tags			Settings
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id for delete Setting"	Format(id)
//	@Success		200	{object}	Response
//	@Router			/v1/settings/{id} [DELETE]
func (h *Handler) SettingDelete(c *gin.Context) {
	inputID := c.Param("id")

	ID, err := strconv.ParseUint(inputID, 10, 32)
	if inputID == "" || err != nil {
		newResponse(c, http.StatusBadRequest, models.ErrNotFoundId.Error())
		logger.Zap.Error("error while get query from url OrderUpdate", logger.Error(err))
		return
	}

	err = h.services.Settings.DeleteByID(c.Request.Context(), uint32(ID))
	if err != nil {
		c.JSON(http.StatusOK, models.EmptyStruct{})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "success"})
}
