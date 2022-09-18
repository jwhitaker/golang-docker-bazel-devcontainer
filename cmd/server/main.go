package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thewhitakers/echoer/pkg/echoer"
)

func main() {
	svc := echoer.NewService()
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		msg := svc.GetDefaultMessage()
		return c.String(http.StatusOK, msg)
	})

	e.POST("/", func(c echo.Context) error {
		requestMsg := c.FormValue("msg")

		msg := svc.GetMessage(string(requestMsg))
		return c.String(http.StatusOK, msg)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
