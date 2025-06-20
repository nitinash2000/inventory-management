package routes

import (
	"inventory-management/handlers"
	"inventory-management/models"
	"inventory-management/repository"
	"inventory-management/services/articles"

	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.Engine) {
	articleMap := map[string]*models.Article{
		"123": {
			ArticleId:   "123",
			ArticleName: "Book",
			Price:       200,
			Stock:       4,
		},
	}

	articleRepo := repository.NewArticleRepo(articleMap)
	articleService := articles.NewArticleService(articleRepo)
	articleHandler := handlers.NewHandler(articleService)

	r.GET("/articles/:id", articleHandler.GetArticle)
	r.POST("/articles", articleHandler.CreateArticle)
	r.DELETE("/articles/:id", articleHandler.DeleteArticle)
	r.PUT("/articles/:id", articleHandler.UpdateArticle)
	r.GET("/articles-list", articleHandler.ListArticles)
	r.PATCH("/articles/:id", articleHandler.UpdateArticleStock)
}
