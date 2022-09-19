package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thewhitakers/echoer/pkg/echoer"
)

func main() {
	svc := echoer.NewService()
	e := echo.New()

	e.GET("/", handleGet(svc))
	e.POST("/", handlePost(svc))

	e.Logger.Fatal(e.Start(":3000"))
}

func handleGet(svc echoer.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		msg := svc.GetDefaultMessage()
		return c.String(http.StatusOK, msg)
	}
}

func handlePost(svc echoer.Service) echo.HandlerFunc {
	type messageRequest struct {
		Message string `json:"message"`
	}

	return func(c echo.Context) error {
		m := new(messageRequest)

		if err := c.Bind(m); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		msg := svc.GetMessage(m.Message)
		return c.String(http.StatusOK, msg)
	}
}
