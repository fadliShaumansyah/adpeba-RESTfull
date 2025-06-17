package main

import (
	"go-adpeba/config"
	"go-adpeba/middleware"
	"go-adpeba/models"
	"go-adpeba/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    dsn := config.GetDSN()
    db, err := config.ConnectDatabase(dsn)

    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    config.DB.AutoMigrate(&models.Leader{})
    config.DB = db

    r := gin.Default()

    r.Use(middleware.ErrorHandler())
    routes.LeaderRoute(r)
    r.Run(":8080")
}
