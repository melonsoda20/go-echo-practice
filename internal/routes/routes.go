package routes

import (
	"ToDoApp/internal/controllers"
	"ToDoApp/internal/handler/hello"

	"github.com/labstack/echo/v4"
)

func GenerateRoutes(echo *echo.Echo) {
	generateHelloRoutes(echo)
	generateToDoRoutes(echo)
}

func generateHelloRoutes(echo *echo.Echo) {
	echo.GET("/hello", hello.GetHello)
}

func generateToDoRoutes(echo *echo.Echo) {
	echo.POST("/todo", controllers.CreateToDo)
}
