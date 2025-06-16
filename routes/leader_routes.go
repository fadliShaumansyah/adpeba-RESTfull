package routes

import (
	"go-adpeba/controllers"

	"github.com/gin-gonic/gin"
)

func LeaderRoute(router *gin.Engine) {
    api := router.Group("/leaders")
    {
        api.GET("", controllers.GetLeaders)
        api.GET("/:id", controllers.GetLeaderByID)
        api.POST("", controllers.CreateLeader)
        api.PUT("/:id", controllers.UpdateLeader)
        api.DELETE("/:id", controllers.DeleteLeader)
    }
}
