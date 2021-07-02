package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Index struct{}

func NewIndex() Handler {
	return &Index{}
}

func (h *Index) GetEndPoint() string {
	return "/index"
}

func (h *Index) GetMethod() string {
	return http.MethodGet
}

func (h *Index) HandlerFunc(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}
