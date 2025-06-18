package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Struct standar untuk semua response
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//
// 📗 SUCCESS RESPONSES
//

// 200 OK
func RespondSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// 201 Created
func RespondCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// 204 No Content (jika tidak perlu kirim body, misalnya delete sukses)
func RespondNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

//
// 📕 ERROR RESPONSES
//

// 400 Bad Request
func RespondBadRequest(c *gin.Context, message string, errors []string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": message,
		"errors": errors,

	})
}

// 401 Unauthorized
func RespondUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		Status:  "error",
		Message: message,
	})
}

// 403 Forbidden
func RespondForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, APIResponse{
		Status:  "error",
		Message: message,
	})
}

// 404 Not Found
func RespondNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		Status:  "error",
		Message: message,
	})
}

// 409 Conflict
func RespondConflict(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, APIResponse{
		Status:  "error",
		Message: message,
	})
}

// 500 Internal Server Error
func RespondInternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		Status:  "error",
		Message: message,
	})
}

// General Error (custom code)
func RespondError(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Status:  "error",
		Message: message,
	})
}

// Mengubah semua pesan validasi ke array string
func FormatValidationError(err error) []string {
	var errors []string

	if verrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range verrs {
			errors = append(errors, fmt.Sprintf("Field '%s' %s", e.Field(), e.Tag()))
		}
	} else {
		// Jika bukan error validasi, tampilkan error langsung
		errors = append(errors, err.Error())
	}

	return errors
}