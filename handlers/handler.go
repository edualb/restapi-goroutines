package handlers

import "github.com/labstack/echo/v4"

type Handler interface {
	GetMethod() string
	GetEndPoint() string
	HandlerFunc(c echo.Context) error
}
