package route

import (
	api2 "echo-sample/internal/api"
	"echo-sample/internal/db"
	"echo-sample/internal/handler"
	"echo-sample/internal/model"
	myMw "echo-sample/middleware"
	echo "github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Debug = true

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// Set Custom MiddleWare
	e.Use(myMw.TransactionHandler(db.NewJsonDB[model.Member]("database.json")))
	e.HTTPErrorHandler = handler.JSONHTTPErrorHandler
	e.GET("/", api2.GetMembers())
	// Routes
	v1 := e.Group("/api/v1")
	v1.POST("/members", api2.PostMember())
	v1.GET("/members", api2.GetMembers())
	v1.GET("/members/:id", api2.GetMember())
	// HOMEWORK: Make this:
	v1.DELETE("/members/:id", api2.DeleteMember())

	v2 := e.Group("/admin")
	// admin group with auth check
	v2.Use(myMw.AlwaysUnauthorized())
	v2.GET("/members", api2.GetMembers())
	return e
}

func toPtr[T any](in T) *T {
	return &in
}
