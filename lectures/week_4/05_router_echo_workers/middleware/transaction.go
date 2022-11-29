package middleware

import (
	"errors"
	"net/http"

	"echo-sample/db"
	"echo-sample/model"

	"github.com/labstack/echo/v4"
)

const (
	TxKey = "Tx"
)

func TransactionHandler(db *db.JsonDB[model.Member]) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			if db == nil {
				return errors.New("db not initialized")
			}

			c.Set(TxKey, db)
			return next(c)
		})
	}
}

func AlwaysUnauthorized() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
		})
	}
}
