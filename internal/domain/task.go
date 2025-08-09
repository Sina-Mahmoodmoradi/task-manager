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
    Update(task *Task) error
    GetByID(id uint) (*Task, error)
    Delete(id uint) error
    GetByUserId(userID uint) ([]Task, error)
}

type TaskUpdate struct {
    Title       *string
    Description *string
    Done        *bool
}

type TaskService interface {
    CreateTask(userID uint, title, description string) (*Task, error)
    GetTaskByID(taskID uint) (*Task, error)
    UpdateTask(taskID uint, userID uint, updates *TaskUpdate) (*Task, error)
    DeleteTask(taskID uint) error
    ListTasksByUser(userID uint) ([]Task, error)
}