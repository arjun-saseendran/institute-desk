package enrollment

import (
	"time"

	"github.com/arjun-saseendran/institute-desk/internal/class"
	"github.com/arjun-saseendran/institute-desk/internal/session"
	"github.com/arjun-saseendran/institute-desk/internal/user"
)

type StatusType string

const (
	Active   StatusType = "ACTIVE"
	Inactive StatusType = "INACTIVE"
)

type Enrollment struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	UserId    uint            `json:"userId"`
	User      user.User       `gorm:"foreignKey:UserId" json:"user,omitempty"`
	ClassId   uint            `json:"classId"`
	Class     class.Class     `gorm:"foreignKey:ClassId" json:"class,omitempty"`
	SessionId uint            `json:"sessionId"`
	Session   session.Session `gorm:"foreignKey:SessionId" json:"session,omitempty"`
	Status    StatusType      `gorm:"default:'ACTIVE'" json:"statusType"`
	StartDate string          `json:"startDate"`
	EndDate   string          `json:"endDate"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
}

func NewEnrollment() *Enrollment {
	return &Enrollment{}
}
