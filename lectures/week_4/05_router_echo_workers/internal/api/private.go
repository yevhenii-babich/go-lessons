package api

import (
	"net/http"
	"strconv"

	"echo-with-workers/internal/db"
	"echo-with-workers/internal/model"

	"github.com/labstack/echo/v4"
)

func DeleteMember() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		number, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		tx := c.Get("Tx").(*db.JsonDB[model.Member])

		tx.Delete(number)
		return c.JSON(http.StatusOK, number)
	}
}
