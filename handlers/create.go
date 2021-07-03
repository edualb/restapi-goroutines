package handlers

import (
	"net/http"

	"github.com/edualb/restapi-goroutines/models"
	"github.com/edualb/restapi-goroutines/services"
	"github.com/labstack/echo/v4"
)

type Create struct {
	service services.Create
}

func NewCreate() Handler {
	return &Create{
		service: services.NewCreate(),
	}
}

func (h *Create) GetEndPoint() string {
	return "/create"
}

func (h *Create) GetMethod() string {
	return http.MethodPost
}

func (h *Create) HandlerFunc(c echo.Context) error {
	var w models.WalkDogService
	if err := c.Bind(&w); err != nil {
		return err
	}
	if err := h.service.Create(&w); err != nil {
		return c.String(http.StatusInternalServerError, "NOT OK")
	}
	return c.String(http.StatusOK, "CREATED")
}
