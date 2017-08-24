package simplyst

// User holds metadata about a book.
type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	UserName  string
}

// UserDatabase provides thread-safe access to a database of books.
type UserDatabase interface {
	Adduser(u *User) (id int64, err error)
	Close()
}
