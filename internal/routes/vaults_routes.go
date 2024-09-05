package routes

import (
	"github.com/Arthur-Conti/keepass-project/internal/controllers"
	"github.com/gin-gonic/gin"
)

func VaultsRoutes(router *gin.RouterGroup) {
	vaultsController := controllers.NewVaultsController()
	router.GET("/:user_id", vaultsController.ListAllVaultsByUserIDHandler)
	router.POST("/", vaultsController.CreateVaultHandler)
}