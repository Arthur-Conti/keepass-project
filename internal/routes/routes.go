package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": true, "status_code": 200})
	})
	userRouter := server.Group("/users")
	UsersRoutes(userRouter)
	vaultsRouter := server.Group("/vaults")
	VaultsRoutes(vaultsRouter)
}