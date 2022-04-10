package controller

import (
	"net/http"
	"redis/service"

	"github.com/labstack/echo"
)

type Controller interface {
	GetHumans(ctx echo.Context) error
	GetHuman(ctx echo.Context) error
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{service: service}
}

func (c *controller) GetHumans(ctx echo.Context) error {
	request := ctx.Request()
	context := request.Context()

	humans, err := c.service.GetHumans(context)
	if err != nil {
		return ctx.NoContent(http.StatusNoContent)
	}
	return ctx.JSON(http.StatusOK, *humans)
}

func (c *controller) GetHuman(ctx echo.Context) error {
	request := ctx.Request()
	context := request.Context()

	humanID := ctx.Param("id")
	human, err := c.service.GetHuman(context, humanID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, *human)
}
