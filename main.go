package main

import (
	"go-adpeba/config"
	"go-adpeba/models"
	"go-adpeba/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDatabase()
    config.DB.AutoMigrate(&models.Leader{})

    r := gin.Default()
    routes.LeaderRoute(r)
    r.Run(":8080")
}
