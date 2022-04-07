package main

import (
	"ToDoApp/internal/extensions/server"
	"ToDoApp/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	server.AddExtensions(e)

	routes.GenerateRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
