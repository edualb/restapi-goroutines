package main

import (
	"net/http"

	"github.com/edualb/restapi-goroutines/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	hs := []handlers.Handler{}
	hs = append(hs, handlers.NewCreate())
	hs = append(hs, handlers.NewIndex())
	hs = append(hs, handlers.NewShow())

	for _, h := range hs {
		switch h.GetMethod() {
		case http.MethodGet:
			e.GET(h.GetEndPoint(), h.HandlerFunc)
		case http.MethodPost:
			e.POST(h.GetEndPoint(), h.HandlerFunc)
		}
	}

	e.Logger.Fatal(e.Start(":8080"))
}
