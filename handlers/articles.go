package handlers

import (
	"inventory-management/dtos"
	"inventory-management/services/articles"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	articleService articles.IArticleService
}

func NewHandler(articleService articles.IArticleService) *Handler {
	return &Handler{
		articleService: articleService,
	}
}

func (h *Handler) GetArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	article, err := h.articleService.GetArticle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func (h *Handler) CreateArticle(ctx *gin.Context) {
	var req *dtos.Article

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.articleService.CreateArticle(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article created successfully"})
}

func (h *Handler) DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.articleService.DeleleArticle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

func (h *Handler) UpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	var req *dtos.Article
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.articleService.UpdateArticle(id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Updated article successfully"})
}

func (h *Handler) ListArticles(ctx *gin.Context) {
	articles, err := h.articleService.ListArticle()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

func (h *Handler) UpdateArticleStock(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *dtos.UpdateStock

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.articleService.UpdateArticleStock(id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article stock updated successfully"})
}
