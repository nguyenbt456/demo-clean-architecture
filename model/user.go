package model

// User represents for users table
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type UserRepository interface {
	GetByID(string) (*User, error)
}

type UserService interface {
	GetByID(string) (*User, error)
}
