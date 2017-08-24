package simplyst

type User struct {
	FirstName string
	LastName  string
	Email     string
	UserName  string
}

type UserDatabase interface {
	Close()
}
