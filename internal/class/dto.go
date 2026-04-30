package class

type InputCreateClass struct {
	Title     string `json:"title" binding:"required,min=2,max=50"`
	Category  string `json:"category" binding:"required,min=2,max=50"`
	Price     int32  `json:"price" binding:"required,min=0"`
	SessionId uint   `json:"sessionId" binding:"required"`
}

type InputUpdateClass struct {
	Title     string `json:"title" binding:"required,min=2,max=50"`
	Category  string `json:"category" binding:"required,min=2,max=50"`
	Price     int32  `json:"price" binding:"required,min=0"`
	SessionId uint   `json:"sessionId" binding:"required"`
}

func NewInputCreateClass() *InputCreateClass {
	return &InputCreateClass{}
}

func NewInputUpdateClass() *InputUpdateClass {
	return &InputUpdateClass{}
}
