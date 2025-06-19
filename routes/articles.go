package routes

import (
	"inventory-management/handlers"

	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.Engine) {
	r.GET("/articles/:id", handlers.GetArticle)
	r.POST("/articles", handlers.CreateArticle)
	r.DELETE("/articles/:id", handlers.DeleteArticle)
	r.PUT("/articles/:id", handlers.UpdateArticle)
	r.GET("/articles-list", handlers.ListArticles)
	r.PATCH("/articles/:id", handlers.UpdateArticleStock)
}
