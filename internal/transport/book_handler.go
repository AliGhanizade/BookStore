package transport

import (
	"BookStore/internal/service"
	"BookStore/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{service: s}
}

func (h *BookHandler) RegisterRoutes(r *gin.Engine) {
	bookGroup := r.Group("/books")
	{
		bookGroup.POST("/create", h.Create)
		bookGroup.GET("/:id", h.GetByID)
		bookGroup.GET("/", h.ListAll)
		bookGroup.PUT("/:id", h.Update)
		bookGroup.DELETE("/:id", h.Delete)
	}
}
func (h *BookHandler) ListAll(ctx *gin.Context) {
	type searchParams struct {
		Title  string `form:"title"`
		Author string `form:"author"`
		Year   uint   `form:"year"`
		Page   int    `form:"page"`
	}
	var params searchParams
	if params.Page == 0 {
		params.Page = 1
	}
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid query parameters")
		return
	}
	if params.Title != "" || params.Author != "" || params.Year != 0 {
		books, err := h.service.Search(params.Title, params.Author, params.Year)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		response.Response(ctx, http.StatusOK, books, "Books retrieved successfully")
		return
	}
	books, err := h.service.Paginate(params.Page, 3)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if len(books) == 0 {
		response.Response(ctx, http.StatusNotFound, nil, "No books found")
		return
	}
	response.Response(ctx, http.StatusOK, books, "Books retrieved successfully")
}

func (h *BookHandler) Create(ctx *gin.Context) {
	var req struct {
		Title  string `json:"title" binding:"required"`
		Author string `json:"author" binding:"required"`
		Year   uint   `json:"year" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book, err := h.service.Create(req.Title, req.Author, req.Year)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Response(ctx, http.StatusCreated, book, "Book created successfully")
}

func (h *BookHandler) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	book, err := h.service.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	response.Response(ctx, http.StatusOK, book, "Book retrieved successfully")
}

func (h *BookHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req struct {
		Title  string `json:"title" binding:"required"`
		Author string `json:"author"`
		Year   uint   `json:"year"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book, err := h.service.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	book.Title = req.Title
	book.Author = req.Author
	book.Year = req.Year
	if err := h.service.Update(book); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Response(ctx, http.StatusOK, book, "Book updated successfully")
}

func (h *BookHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	book, err := h.service.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	if err := h.service.Delete(book); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Response(ctx, http.StatusOK, nil, "Book deleted successfully")
}
