package controllers

import (
	"ToDoApp/internal/models/requests"
	"ToDoApp/internal/services/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateToDo(c echo.Context) (err error) {
	r := new(*requests.CreateToDo)
	if err = c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewValidatorError(err))
	}

	if err := c.Validate(r); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewValidatorError(err))
	}

	return c.JSON(http.StatusOK, r)
}
