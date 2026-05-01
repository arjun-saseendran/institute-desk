package enrollment

type InputCreateEnrollment struct {
	UserId    uint `json:"userId" binding:"required"`
	ClassId   uint `json:"classId" binding:"required"`
	SessionId uint `json:"sessionId" binding:"required"`
}
