package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName   string
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{"api/user", userService}

}

func (handler *UserHandler) RegisterEndPoints(r *gin.Engine) {
	userGroup := r.Group(handler.groupName)

	userGroup.POST("", handler.CreateUser)
	userGroup.GET("", handler.ListUsers)

}

func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	userData := NewInputUser()
	err := ctx.BindJSON(&userData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to bind user data"})
		return
	}
	newUser, err := handler.userService.Create(userData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to create new user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "user created successfully", "data": newUser})
}

func (handler *UserHandler) ListUsers(ctx *gin.Context) {
	allUsers, err := handler.userService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to get user data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "user data fetched successfully", "data": allUsers})
}
