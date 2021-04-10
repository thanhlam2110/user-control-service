package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/thanhlam/user-control-service/service"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "SSO SERVICE")
	})
	//<---------------------------SSO------------------------------>
	e.POST("/api/user/userOrderCommand", service.UserOrderCommand)
	e.Logger.Fatal(e.Start(":1323"))
}
