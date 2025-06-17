package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct standar untuk response
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Response sukses
func RespondSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Response error
func RespondError(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Status:  "error",
		Message: message,
	})
}
