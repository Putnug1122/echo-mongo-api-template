package main

import (
	"echo-test/configs"
	"echo-test/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDB()
	routes.UserRoute(e)
	e.Logger.Fatal(e.Start(":6969"))
}
