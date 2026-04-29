package session

import "time"

type Session struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"name"`
	StartTime string    `json:"startTime"`
	EndTime   string    `json:"endTime"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func NewSession() *Session {
	return &Session{}
}
