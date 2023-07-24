package router

import (
	infra "go-migratre-sample/infra/postgres"
	"go-migratre-sample/interface/handler"
	"go-migratre-sample/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserDIRouting(db *gorm.DB, e *echo.Echo) {

	userRepository := infra.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	e.POST("/users", userHandler.CreateUser())
	e.PUT("/users/:id", userHandler.UpdateUser())
	e.DELETE("/users/:id", userHandler.DeleteUser())
	e.GET("/users", userHandler.GetUserList())
	e.GET("/users/:id", userHandler.GetUserOne())
}
