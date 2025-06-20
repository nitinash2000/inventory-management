package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(r *gin.Engine, db *gorm.DB) {
	ArticleRoutes(r, db)
}
