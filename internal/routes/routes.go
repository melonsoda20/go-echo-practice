package routes

import (
	"ToDoApp/internal/handler/hello"

	"github.com/labstack/echo/v4"
)

func GenerateRoutes(echo *echo.Echo) {
	generateGetRoutes(echo)
}

func generateGetRoutes(echo *echo.Echo) {
	echo.GET("/hello", hello.GetHello)
}
