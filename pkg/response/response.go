// pkg/response/response.go
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status  string      `json:"status"`            
	Data    interface{} `json:"data,omitempty"`    
	Message string      `json:"message,omitempty"` 
}
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Status: "success",
		Data:   data,
	})
}

func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, ErrorResponse{
		Status:  "error",
		Message: message,
	})
}
