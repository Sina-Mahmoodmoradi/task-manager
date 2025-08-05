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



type TaskService interface {
    CreateTask(userID uint, title, description string) (*Task, error)
    GetTaskByID(taskID uint) (*Task, error)
    UpdateTask(taskID uint, title, description string) (*Task, error)
    DeleteTask(taskID uint) error
    ListTasksByUser(userID uint) ([]Task, error)
}