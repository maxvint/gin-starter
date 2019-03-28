package model

type Model struct {
	Model interface{}
}

type Task struct {
	UserId  int32
	Title   string
	Summary string
}

func (s *Task) TableName() string {
	return "task"
}
