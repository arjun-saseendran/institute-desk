package user

import "time"

type UserType string

const (
	Admin   UserType = "ADMIN"
	Student UserType = "STUDENT"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Mobile    string    `gorm:"uniqueIndex" json:"mobile"`
	Role      UserType  `gorm:"default:'STUDENT'" json:"userType"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func NewUser() *User {
	return &User{}
}
