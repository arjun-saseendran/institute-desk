package user

type InputCreateUser struct {
	Name    string `json:"name" binding:"required,min=2,max=50"`
	Address string `json:"address" binding:"required"`
	Mobile  string `json:"mobile" binding:"required,min=10,max=15"`
}

type InputUpdateUser struct {
	Name    string `json:"name" binding:"required,min=2,max=50"`
	Address string `json:"address" binding:"required"`
	Mobile  string `json:"mobile" binding:"required,min=10,max=15"`
}

func NewInputCreateUser() *InputCreateUser {
	return &InputCreateUser{}
}

func NewInputUpdateUser() *InputUpdateUser {
	return &InputUpdateUser{}
} 
