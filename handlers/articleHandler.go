package handlers

import (
	"inventory-management/dtos"
	"inventory-management/services/articles"
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleHandler struct {
	articleService articles.ArticleService
}

func NewArticleHandler(articleService articles.ArticleService) *articleHandler {
	return &articleHandler{
		articleService: articleService,
	}
}

func (a *articleHandler) GetArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	article, err := a.articleService.GetArticle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func (a *articleHandler) CreateArticle(ctx *gin.Context) {
	var req *dtos.Article

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = a.articleService.CreateArticle(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article created successfully"})
}

func (a *articleHandler) DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	err := a.articleService.DeleteArticle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

func (a *articleHandler) UpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dtos.Article
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = a.articleService.UpdateArticle(id, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Updated article successfully"})
}

func (a *articleHandler) ListArticles(ctx *gin.Context) {
	articles, err := a.articleService.ListArticle()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

func (a *articleHandler) UpdateArticleStock(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *dtos.UpdateStock

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = a.articleService.UpdateArticleStock(id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article stock updated successfully"})
}
