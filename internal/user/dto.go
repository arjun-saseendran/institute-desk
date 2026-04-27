package user

type InputUser struct {
	Name    string `json:"name" binding:"required,min=2,max=50"`
	Address string `json:"address" binding:"required"`
	Mobile  string `json:"mobile" binding:"required,min=10,max=15"`
}

type UpdateUser struct {
	Name    string `json:"name" binding:"required,min=2,max=50"`
	Address string `json:"address" binding:"required"`
	Mobile  string `json:"mobile" binding:"required,min=10,max=15"`
}

func NewInputUser() *InputUser {
	return &InputUser{}
}

func NewUpdateUser() *UpdateUser {
	return &UpdateUser{}
} 
