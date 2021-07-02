package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Create struct{}

func NewCreate() Handler {
	return &Create{}
}

func (h *Create) GetEndPoint() string {
	return "/create"
}

func (h *Create) GetMethod() string {
	return http.MethodPost
}

func (h *Create) HandlerFunc(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}
