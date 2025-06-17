package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware recover panic + log error
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("[PANIC] %v\n", rec)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": "Internal server error",
				})
			}
		}()
		c.Next()

		// Tangani error dari c.Errors
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				log.Printf("[ERROR] %v\n", e.Err)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": c.Errors[0].Error(),
			})
			c.Abort()
			return
		}
	}
}
