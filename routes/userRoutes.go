package routes

import (
	"inventory-management/handlers"
	"inventory-management/repository"
	"inventory-management/services/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepo(db)
	addressRepo := repository.NewAddressRepo(db)

	userService := users.NewUserService(userRepo, addressRepo)
	userHandler := handlers.NewUserHandler(userService)

	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.CreateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
}
