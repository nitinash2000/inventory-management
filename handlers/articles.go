package handlers

import (
	"inventory-management/dtos"
	"inventory-management/services/articles"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	a := articles.NewArticleService()
	article, err := a.GetArticle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func CreateArticle(ctx *gin.Context) {
	var req *dtos.Article

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	a := articles.NewArticleService()
	err = a.CreateArticle(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article created successfully"})
}

func DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	a := articles.NewArticleService()
	err := a.DeleleArticle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

func UpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	var req *dtos.Article
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	a := articles.NewArticleService()
	err = a.UpdateArticle(id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Updated article successfully"})
}

func ListArticles(ctx *gin.Context) {
	a := articles.NewArticleService()
	articles, err := a.ListArticle()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

func UpdateArticleStock(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *dtos.UpdateStock

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	a := articles.NewArticleService()
	err = a.UpdateArticleStock(id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Article stock updated successfully"})
}
