package routes

import (
	"inventory-management/handlers"
	"inventory-management/repository"
	"inventory-management/services/articles"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArticleRoutes(r *gin.Engine, db *gorm.DB) {
	articleRepo := repository.NewArticleRepo(db)
	articleService := articles.NewArticleService(articleRepo)
	articleHandler := handlers.NewHandler(articleService)

	r.GET("/articles/:id", articleHandler.GetArticle)
	r.POST("/articles", articleHandler.CreateArticle)
	r.DELETE("/articles/:id", articleHandler.DeleteArticle)
	r.PUT("/articles/:id", articleHandler.UpdateArticle)
	r.GET("/articles-list", articleHandler.ListArticles)
	r.PATCH("/articles/:id", articleHandler.UpdateArticleStock)
}
