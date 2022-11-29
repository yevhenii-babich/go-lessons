package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg, ok = he.Message.(string)
	}
	if !c.Response().Committed {
		_ = c.JSON(code, map[string]interface{}{
			"statusCode": code,
			"message":    msg,
		})
	}
}
