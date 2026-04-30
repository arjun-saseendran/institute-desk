package class

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	groupName    string
	classService ClassService
}

func NewClassHandler(classService ClassService) *ClassHandler {
	return &ClassHandler{"api/class", classService}
}

func (handler *ClassHandler) RegisterEndPoints(r *gin.Engine) {
	classGroup := r.Group(handler.groupName)

	classGroup.POST("/", handler.CreateClass)
	classGroup.GET("/", handler.GetClasses)
	classGroup.GET("/", handler.GetClass)
	classGroup.PUT("/", handler.UpdateClass)
	classGroup.DELETE("/", handler.DeleteClass)
}

func (handler *ClassHandler) CreateClass(ctx *gin.Context) {
	classData := NewInputCreateClass()
	err := ctx.BindJSON(&classData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to bind class data"})
		return
	}
	newClass, err := handler.classService.CreateClass(classData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to create new class"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "class created", "data": newClass})
}

func (handler *ClassHandler) GetClasses(ctx *gin.Context) {
	allClasses, err := handler.classService.GetClasses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed fetch class data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "class data fetched successfully", "data": allClasses})
}

func (handler *ClassHandler) GetClass(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid class id"})
		return
	}
	singleClass, err := handler.classService.GetClass(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to class data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "class data fetched successfully", "data": singleClass})
}

func (handler *ClassHandler) UpdateClass(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid class id"})
		return
	}
	classData := NewInputUpdateClass()

	err := ctx.BindJSON(&classData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to bind class data"})
		return
	}
	updatedClassData, err := handler.classService.UpdateClass(id, classData)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to update class data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "class data updated successfully", "data": updatedClassData})

}

func (handler *ClassHandler) DeleteClass(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid class id"})
		return
	}
	err := handler.classService.DeleteClass(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to delete class data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "class deleted successfully"})
}
