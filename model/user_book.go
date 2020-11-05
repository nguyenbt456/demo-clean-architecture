package model

// UserBook represents for user_books table
type UserBook struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	BookID string `json:"book_id"`
}

type UserBookRepository interface {
	GetByUserID(string) ([]*UserBook, error)
}

type UserBookService interface {
	GetByUserID(string) ([]*UserBook, error)
}
