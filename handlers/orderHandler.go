package handlers

import (
	"inventory-management/dtos"
	"inventory-management/services/orders"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService orders.OrderService
}

func NewOrderHandler(orderService orders.OrderService) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (o *orderHandler) GetOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	order, err := o.orderService.GetOrder(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (o *orderHandler) CreateOrder(ctx *gin.Context) {
	var req *dtos.Order

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = o.orderService.CreateOrder(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

func (o *orderHandler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	err := o.orderService.DeleleOrder(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func (o *orderHandler) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dtos.Order
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = o.orderService.UpdateOrder(id, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Updated order successfully"})
}
