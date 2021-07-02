package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Show struct{}

func NewShow() Handler {
	return &Show{}
}

func (h *Show) GetEndPoint() string {
	return "/show"
}

func (h *Show) GetMethod() string {
	return http.MethodGet
}

func (h *Show) HandlerFunc(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}
