package class

import (
	"time"

	"github.com/arjun-saseendran/institute-desk/internal/session"
)

type Class struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	Title     string          `json:"title"`
	Category  string          `json:"category"`
	Price     int32           `json:"price"`
	SessionId uint            `json:"sessionId"`
	Session   session.Session `gorm:"foreignKey:SessionId" json:"session,omitempty"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
}

func NewClass() *Class {
	return &Class{}
}
