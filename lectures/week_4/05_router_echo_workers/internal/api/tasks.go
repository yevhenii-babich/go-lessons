package api

import (
	"net/http"

	"echo-with-workers/internal/workerpool"

	"github.com/labstack/echo/v4"

	"echo-with-workers/internal/model"
)

func PostTask() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		m := model.NewTask(c.Logger())
		if err := c.Bind(&m); err != nil {
			c.Logger().Errorf("can't bind: %v", err)
			return err
		}

		pool, ok := c.Get("Wrk").(*workerpool.Pool)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				Code:    500,
				Message: "context data error",
			})
		}
		pool.Exec(m, m.InParam)

		return c.JSON(http.StatusCreated, m)
	}
}
