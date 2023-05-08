package delivery

import (
	"net/http"
	"strconv"
	"time"
	entryUsecase "timetracker/internal/Entry/usecase"
	"timetracker/models"
	"timetracker/models/dto"
	"timetracker/pkg"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Delivery struct {
	EntryUC entryUsecase.UsecaseI
}

// CreateEntry godoc
// @Summary      Create entry
// @Description  Create entry
// @Tags     	 entry
// @Accept	 application/json
// @Produce  application/json
// @Param    entry body dto.ReqCreateUpdateEntry true "entry info"
// @Success  200 {object} pkg.Response{body=dto.RespEntry} "success update entry"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 400 {object} echo.HTTPError "bad req"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /entry/create [post]
func (delivery *Delivery) CreateEntry(c echo.Context) error {

	var reqEntry dto.ReqCreateUpdateEntry
	err := c.Bind(&reqEntry)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqEntry); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error("can't parse context user_id")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	entry := reqEntry.ToModelEntry()
	entry.UserID = &userId
	err = delivery.EntryUC.CreateEntry(entry)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respEntry := dto.GetResponseFromModelEntry(entry)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respEntry})
}

// GetEntry godoc
// @Summary      Show a post
// @Description  Get entry by id
// @Tags     	 entry
// @Accept	 application/json
// @Produce  application/json
// @Param id  path int  true  "Entry ID"
// @Success  200 {object} pkg.Response{body=dto.RespEntry} "success get entry"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /entry/{id} [get]
func (delivery *Delivery) GetEntry(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}
	entry, err := delivery.EntryUC.GetEntry(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respEntry := dto.GetResponseFromModelEntry(entry)
	return c.JSON(http.StatusOK, pkg.Response{Body: *respEntry})
}

// UpdateEntry godoc
// @Summary      Update an entry
// @Description  Update an entry
// @Tags     	 entry
// @Accept	 application/json
// @Produce  application/json
// @Param    entry body dto.ReqCreateUpdateEntry true "entry info"
// @Success  200 {object} pkg.Response{body=dto.RespEntry} "success update entry"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /entry/edit [post]
func (delivery *Delivery) UpdateEntry(c echo.Context) error {

	var reqEntry dto.ReqCreateUpdateEntry
	err := c.Bind(&reqEntry)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqEntry); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error("can't parse context user_id")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	entry := reqEntry.ToModelEntry()
	entry.UserID = &userId
	err = delivery.EntryUC.UpdateEntry(entry)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respEntry := dto.GetResponseFromModelEntry(entry)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respEntry})
}

// DeleteEntry godoc
// @Summary      Delete an entry
// @Description  Delete an entry
// @Tags     	 entry
// @Accept	 application/json
// @Param id path int  true  "Entry ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 404 {object} echo.HTTPError "can't find entry with such id"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /entry/{id} [delete]
func (delivery *Delivery) DeleteEntry(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error("can't parse context user_id")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	err = delivery.EntryUC.DeleteEntry(id, userId)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

// GetUserEntries godoc
// @Summary      Get user entries
// @Description  Get user entries or get user entries for a day
// @Tags     entry
// @Produce  application/json
// @Param        day    query     string  false  "day for events"
// @Success  200 {object} pkg.Response{body=[]dto.RespEntry} "success get entries"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /entry/usr [get]
func (delivery *Delivery) GetUserEntries(c echo.Context) error {

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error("can't parse context user_id")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	day := c.QueryParam("day")
	var entries []*models.Entry
	var err error

	if day == "" {
		entries, err = delivery.EntryUC.GetUserEntries(userId)
	} else {
		date, err := time.Parse("2006-01-02", day)

		if err != nil {
			c.Logger().Error("invalid data format, should be YYYY-MM-DD")
			return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
		}

		entries, err = delivery.EntryUC.GetUserEntriesForDay(userId, date)
	}

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respEnties := dto.GetResponseFromModelEntries(entries)

	return c.JSON(http.StatusOK, pkg.Response{Body: respEnties})
}

func handleError(err error) *echo.HTTPError {
	causeErr := errors.Cause(err)
	switch {
	case errors.Is(causeErr, models.ErrNotFound):
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
	case errors.Is(causeErr, models.ErrBadRequest):
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	case errors.Is(causeErr, models.ErrPermissionDenied):
		return echo.NewHTTPError(http.StatusForbidden, models.ErrPermissionDenied.Error())
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
	}
}

func NewDelivery(e *echo.Echo, eu entryUsecase.UsecaseI) {
	handler := &Delivery{
		EntryUC: eu,
	}

	e.POST("/entry/create", handler.CreateEntry)
	e.POST("/entry/edit", handler.UpdateEntry)
	e.GET("/entry/:id", handler.GetEntry)
	e.DELETE("/entry/:id", handler.DeleteEntry)
	e.GET("/entry/usr", handler.GetUserEntries)
}
