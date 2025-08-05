package domain

type User struct {
    ID       uint
    Username string
    Password string
    Tasks    []Task
}

type UserRepository interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
}