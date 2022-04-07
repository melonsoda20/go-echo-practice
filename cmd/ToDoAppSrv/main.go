package main

import (
	"ToDoApp/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.GenerateRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
