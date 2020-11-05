package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/demo-clean-architecture/controller"
	"github.com/nguyenbt456/demo-clean-architecture/database"
	"github.com/nguyenbt456/demo-clean-architecture/repository"
	"github.com/nguyenbt456/demo-clean-architecture/service"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(database.GetDB())
	userService := service.NewUserService(userRepo)
	userHandler := controller.NewUserHandler(userService)

	userBookRepo := repository.NewUserBookRepository(database.GetDB())

	bookRepo := repository.NewBookRepository(database.GetDB())
	bookService := service.NewBookService(bookRepo, userBookRepo, userRepo)
	bookHandler := controller.NewBookHandler(bookService)

	router.GET("/users/:user_id", userHandler.GetByID)
	router.GET("/users/:user_id/books", bookHandler.GetByUserID)

	return router
}
