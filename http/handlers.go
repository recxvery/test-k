package http

import (
	"errors"
	"kaspersky-testovoe/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ThreatHandler struct {
	repo db.DBRepository
}

func NewThreatHandler(db_conn db.DBRepository) *ThreatHandler {
	return &ThreatHandler{
		repo: db_conn,
	}
}

func (h *ThreatHandler) GetThreatHandler(c echo.Context) error {
	md5Param := c.Param("md5")

	ctx := c.Request().Context()

	usr, err := h.repo.GetThreat(ctx, md5Param)
	if err != nil {
		//fix errors logging
		if errors.Is(err, db.ErrThreatNotFound) {
			return echo.NewHTTPError(
				http.StatusNotFound,
				"There is no such Threat",
			)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *ThreatHandler) InsertThreatHandler(c echo.Context) error {
	var ThreatInsertRequest ThreatCreateRequest

	if err := c.Bind(&ThreatInsertRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect JSON format")
	}

	newThreat, err := ValidateToCreate(ThreatInsertRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	threat, err := h.repo.InsertThreat(ctx, newThreat)
	if err != nil {
		if errors.Is(err, db.ErrThreatAlreadyExists) {
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Can't create Treat")
	}

	return c.JSON(http.StatusCreated, threat)
}

func (h *ThreatHandler) UpdateThreatHandler(c echo.Context) error {
	md5 := c.Param("md5")
	var threatUpdReq ThreatUpdateRequest

	if err := c.Bind(&threatUpdReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect JSON format")
	}

	ctx := c.Request().Context()

	threatDescripiton, err := ValidteThreatUpdate(threatUpdReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	threat, err := h.repo.UpdateThreat(ctx, md5, threatDescripiton)
	if err != nil {
		if errors.Is(err, db.ErrThreatNotFound) {
			return echo.NewHTTPError(
				http.StatusNotFound,
				err.Error(),
			)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, threat)
}
