package main

import (
	"ToDoApp/internal/handler/hello"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", hello.GetHello)
	e.Logger.Fatal(e.Start(":1323"))
}
