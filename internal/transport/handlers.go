package transport

import (
	"errors"
	"net/http"

	"kaspersky-testovoe/internal/core"
	"kaspersky-testovoe/internal/db"
	"kaspersky-testovoe/internal/service"

	"github.com/labstack/echo/v4"
)

type ThreatHandler struct {
	service *service.ThreatService
}

func NewThreatHandler(svc *service.ThreatService) *ThreatHandler {
	return &ThreatHandler{
		service: svc,
	}
}

func (h *ThreatHandler) GetThreatHandler(c echo.Context) error {
	md5Param := c.Param("md5")
	ctx := c.Request().Context()

	threat, err := h.service.GetThreat(ctx, md5Param)
	if err != nil {
		switch {
		case errors.Is(err, core.IncorrectMD5Length),
			errors.Is(err, core.IncorrectMD5Format):
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())

		case errors.Is(err, db.ErrThreatNotFound):
			return echo.NewHTTPError(http.StatusNotFound, "There is no such Threat")

		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	return c.JSON(http.StatusOK, threat)
}

func (h *ThreatHandler) InsertThreatHandler(c echo.Context) error {
	var request ThreatCreateRequest

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect JSON format")
	}

	// DTO -> DB model
	threatModel := ToThreatModel(request)

	ctx := c.Request().Context()

	threat, err := h.service.CreateThreat(ctx, threatModel)
	if err != nil {
		switch {
		case errors.Is(err, core.IncorrectMD5Length),
			errors.Is(err, core.IncorrectMD5Format),
			errors.Is(err, core.IncorrectName),
			errors.Is(err, core.IncorrectDescription),
			errors.Is(err, core.IncorrectSeverity),
			errors.Is(err, core.IncorrectThreatType):
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())

		case errors.Is(err, db.ErrThreatAlreadyExists):
			return echo.NewHTTPError(http.StatusConflict, err.Error())

		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Can't create Threat")
		}
	}

	return c.JSON(http.StatusCreated, threat)
}

func (h *ThreatHandler) UpdateThreatHandler(c echo.Context) error {
	md5 := c.Param("md5")

	var request ThreatUpdateRequest

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect JSON format")
	}

	ctx := c.Request().Context()

	threat, err := h.service.UpdateThreat(ctx, md5, request.Name)
	if err != nil {
		switch {
		case errors.Is(err, core.IncorrectMD5Length),
			errors.Is(err, core.IncorrectMD5Format),
			errors.Is(err, core.IncorrectName):
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())

		case errors.Is(err, db.ErrThreatNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())

		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}
	}

	return c.JSON(http.StatusOK, threat)
}
