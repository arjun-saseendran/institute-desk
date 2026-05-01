package enrollment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	groupName         string
	enrollmentService EnrollmentService
}

func NewEnrollmetHandler(enrollmentService EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{"api/enrollment", enrollmentService}
}

func (handler *EnrollmentHandler) RegisterEndPoints(r *gin.Engine) {
	enrollmentGroup := r.Group(handler.groupName)

	enrollmentGroup.POST("", handler.CreateEnrollment)
	enrollmentGroup.GET("", handler.GetEnrollments)
	enrollmentGroup.GET("/:id", handler.GetEnrollment)
	enrollmentGroup.DELETE("/:id", handler.DeleteEnrollment)

}

func (handler *EnrollmentHandler) CreateEnrollment(ctx *gin.Context) {
	enrolllmentData := NewInputCreateEnrollment()
	err := ctx.BindJSON(&enrolllmentData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to bind enrollment data"})
		return
	}

	newEnrollment, err := handler.enrollmentService.CreateEnrollment(enrolllmentData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to create enrollment"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "enrollment created successfully", "data": newEnrollment})
}

func (handler *EnrollmentHandler) GetEnrollments(ctx *gin.Context) {
	allEnrolllments, err := handler.enrollmentService.GetEnrollments()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to get enrollment data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "enrollment data fetched successfully", "data": allEnrolllments})
}

func (handler *EnrollmentHandler) GetEnrollment(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid enrollment id"})
		return
	}
	enrollment, err := handler.enrollmentService.GetEnrollment(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to get enrollment data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "enrollment data fetched successfully", "data": enrollment})
}

func (handler *EnrollmentHandler) DeleteEnrollment(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid enrollment id"})
		return
	}
	err := handler.enrollmentService.DeleteEnrollment(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to delete enrollment"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "enrollment deleted successfully"})

}
