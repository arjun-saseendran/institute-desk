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
	userGroup.GET("", handler.GetUsers)
	userGroup.GET("/:id", handler.GetUser)

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

func (handler *UserHandler) GetUsers(ctx *gin.Context) {
	allUsers, err := handler.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to get user data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "user data fetched successfully", "data": allUsers})
}

func (handler *UserHandler) GetUser(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid user id"})
		return
	}
	singleUser, err := handler.userService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to get user data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "user data fetched successfully", "data": singleUser})
}
