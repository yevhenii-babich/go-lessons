package route

import (
	api2 "echo-with-workers/internal/api"
	"echo-with-workers/internal/db"
	"echo-with-workers/internal/handler"
	"echo-with-workers/internal/model"
	"echo-with-workers/internal/workerpool"
	myMw "echo-with-workers/middleware"
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
	e.Use(myMw.WorkerHandler(workerpool.NewPool(3)))
	e.HTTPErrorHandler = handler.JSONHTTPErrorHandler
	e.GET("/", api2.GetMembers())
	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/members", api2.PostMember())
		v1.GET("/members", api2.GetMembers())
		v1.GET("/members/:id", api2.GetMember())
		// HOMEWORK: Make this:
		v1.DELETE("/members/:id", api2.DeleteMember())
		v1.POST("/task", api2.PostTask())

	}

	g3 := e.Group("/tasks")
	{
		g3.POST("/add", api2.PostTask())
	}

	v2 := e.Group("/admin")
	{
		// admin group with auth check
		v2.Use(myMw.AlwaysUnauthorized())
		v2.GET("/members", api2.GetMembers())

	}

	return e
}

func toPtr[T any](in T) *T {
	return &in
}
