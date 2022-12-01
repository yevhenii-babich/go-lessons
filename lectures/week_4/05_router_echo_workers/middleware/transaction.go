package middleware

import (
	"errors"
	"net/http"

	"echo-with-workers/internal/workerpool"

	"echo-with-workers/internal/db"
	"echo-with-workers/internal/model"

	"github.com/labstack/echo/v4"
)

const (
	TxKey     = "Tx"
	WorkerKey = "Wrk"
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

func WorkerHandler(pool *workerpool.Pool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			if pool == nil {
				return errors.New("pool not initialized")
			}

			c.Set(WorkerKey, pool)
			return next(c)
		})
	}
}
