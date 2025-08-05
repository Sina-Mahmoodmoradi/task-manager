package domain

type Task struct {
    ID     uint
    Title  string
    Description string
    Done   bool
    UserID uint
}


type TaskRepository interface {
    Create(task *Task) error
    GetByID(id uint) (*Task, error)
}