package routes

import (
	"inventory-management/handlers"
	"inventory-management/repository"
	"inventory-management/services/orders"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(r *gin.Engine, db *gorm.DB) {
	orderRepo := repository.NewOrderRepo(db)
	orderItemRepo := repository.NewOrderItemRepo(db)

	orderService := orders.NewOrderService(orderRepo, orderItemRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	r.GET("/orders/:id", orderHandler.GetOrder)
	r.POST("/orders", orderHandler.CreateOrder)
	r.DELETE("/orders/:id", orderHandler.DeleteOrder)
	r.PUT("/orders/:id", orderHandler.UpdateOrder)
}
