package routes

import (
	"github.com/Arthur-Conti/keepass-project/internal/controllers"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.RouterGroup) {
	usersController := controllers.NewUserController()
	router.GET("/", usersController.ListAllUsersHandler)
	router.GET("/:id", usersController.GetUserByIDHandler)
	router.POST("/", usersController.CreateUserHandler)
}