package model

type (
	Task struct {
		ID          int
		Name        string
		Description string
		Project_id  int
		User_id     int
		Status      string
	}
)
