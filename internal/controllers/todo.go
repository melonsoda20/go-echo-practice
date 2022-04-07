package controllers

import (
	"ToDoApp/internal/models/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateToDo(c echo.Context) (err error) {
	r := new(requests.CreateToDo)
	if err = c.Bind(r); err != nil {
		return
	}

	return c.JSON(http.StatusOK, r)
}
