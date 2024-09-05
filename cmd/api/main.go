package main

import (
	"github.com/Arthur-Conti/keepass-project/internal/config"
	"github.com/Arthur-Conti/keepass-project/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfigs()
	server := gin.Default()
	routes.RegisterRoutes(server)
	port := "3000"
	server.Run(":" + port)
}