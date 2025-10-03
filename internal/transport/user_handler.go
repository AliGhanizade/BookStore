// internal/transport/user_handler.go
package transport

import (
	"BookStore/pkg/response"
	"BookStore/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", h.Register)
		userGroup.POST("/login", h.Login)
		userGroup.GET("/:id", h.GetByID)
		userGroup.GET("/", h.ListAll)
		userGroup.PUT("/:id", h.Update)
		userGroup.DELETE("/:id", h.Delete)
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.service.Register(req.Name, req.Phone, req.Password)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Response(ctx, http.StatusOK, user, "User registered successfully")
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.service.Login(req.Phone, req.Password)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	response.Response(ctx, http.StatusOK, user, "User logged in successfully")
}

func (h *UserHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := h.service.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.Response(ctx, http.StatusOK, user, "User retrieved successfully")
}

func (h *UserHandler) ListAll(ctx *gin.Context) {
	users, err := h.service.ListAll()
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Response(ctx, http.StatusOK, users, "users retrieved successfully")
}

func (h *UserHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req struct {
		Name       string `json:"name"`
		Phone      string `json:"phone"`
		Password   string `json:"password"`
		Permission string `json:"permission"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.service.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	user.Name = req.Name
	user.Phone = req.Phone
	user.Password = req.Password
	user.Permission = req.Permission

	if err := h.service.Update(user); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Response(ctx, http.StatusOK, user, "user updated successfully")
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := h.service.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	if err := h.service.Delete(user); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Response(ctx, http.StatusOK, nil, "user deleted successfully")
}
