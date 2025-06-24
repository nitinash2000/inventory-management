package handlers

import (
	"inventory-management/dtos"
	"inventory-management/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.UserService
}

func NewUserHandler(userService users.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (c *userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.userService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *userHandler) CreateUser(ctx *gin.Context) {
	var req *dtos.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.userService.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (c *userHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (c *userHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dtos.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.userService.UpdateUser(id, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Updated user successfully"})
}
