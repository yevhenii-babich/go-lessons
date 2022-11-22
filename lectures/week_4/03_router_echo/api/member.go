package api

import (
	"net/http"
	"strconv"

	"echo-sample/db"
	"echo-sample/model"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo/v4"
)

func PostMember() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		m := new(model.Member)
		if err := c.Bind(&m); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		tx := c.Get("Tx").(*db.JsonDB[model.Member])

		member := model.NewMember(m.Number, m.Name, m.Position)

		if err := member.Save(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusCreated, member)
	}
}

func GetMember() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		number, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		tx := c.Get("Tx").(*db.JsonDB[model.Member])

		member := new(model.Member)
		if err := member.Load(tx, number); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound, "Member does not exists.")
		}
		return c.JSON(http.StatusOK, member)
	}
}

func GetMembers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*db.JsonDB[model.Member])

		position := c.QueryParam("position")
		members := new(model.Members)
		if err = members.Load(tx, position); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound, "Member does not exists.")
		}

		return c.JSON(http.StatusOK, members)
	}
}
