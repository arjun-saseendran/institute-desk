package session

type InputCreateSession struct {
	Title     string `json:"title" binding:"required,min=2,max=50"`
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime" bindgin:"required"`
}

type InputUpdateSession struct {
	Title     string `json:"title" binding:"required,min=2,max=50"`
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
}

func NewInputCreateSession() *InputCreateSession {
	return &InputCreateSession{}
}

func NewInputUpdateSession() *InputUpdateSession {
	return &InputUpdateSession{}
}
