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


type UserService interface {
    Register(username, password string) (*User, error)
    Login(username, password string) (string, error) // returns JWT token (for example)

    GetByID(id uint) (*User, error)
    GetByUsername(username string) (*User, error)
    List() ([]User, error)
}
