package route

import (
	"echo-sample/api"
	"echo-sample/db"
	"echo-sample/handler"
	myMw "echo-sample/middleware"
	"echo-sample/model"
	echo "github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

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
	e.GET("/", api.GetMembers())
	// Routes
	v1 := e.Group("/api/v1")
	v1.POST("/members", api.PostMember())
	v1.GET("/members", api.GetMembers())
	v1.GET("/members/:id", api.GetMember())
	// HOMEWORK: Make this:
	// v1.DELETE("/members/:id", api.DeleteMember())

	v2 := e.Group("/admin")
	// admin group with auth check
	v2.Use(myMw.AlwaysUnauthorized())
	v2.GET("/members", api.GetMembers())
	return e
}
