package domain

type Task struct {
    ID     uint
    Title  string
    Description string
    Done   bool
    UserID uint
}
