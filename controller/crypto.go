package controller

import (
	"cryptogo/common"
	"cryptogo/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	CryptoController struct {
	}
	TokenRequest struct {
		ID       string `query:"id"`
		Currency string `query:"currency"`
	}
)

func (controller CryptoController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.GET,
			Path:    "/crypto",
			Handler: controller.List,
		},
	}
}

func (controller CryptoController) List(c echo.Context) error {
	tokenRequest := new(TokenRequest)
	if err := c.Bind(tokenRequest); err != nil {
		return err
	}
	token := service.GetCryptoService().GetListToken(tokenRequest.ID, tokenRequest.Currency)
	return c.JSON(http.StatusOK, token)
}
