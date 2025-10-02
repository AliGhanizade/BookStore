// pkg/response/response.go
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(ctx *gin.Context, code int, data interface{}, message string) {
	ctx.JSON(http.StatusOK, response{
		Code:    code,
		Data:   data,
		Message: message,
	})
}


func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, errorResponse{
		Code:    code,
		Message: message,
	})
}
