package model

// Book represents for books table
type Book struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BookRepository interface {
	GetByID(string) (*Book, error)
}

type BookService interface {
	GetByID(string) (*Book, error)
	GetByUserID(string) ([]*Book, error)
}
