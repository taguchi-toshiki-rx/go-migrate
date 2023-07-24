package main

import (
	infra "go-migratre-sample/infra/postgres"
	"go-migratre-sample/interface/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	db, err := infra.OpenDB()
	if err != nil {
		panic(err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	router.UserDIRouting(db, e)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
