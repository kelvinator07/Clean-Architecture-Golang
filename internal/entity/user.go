package entity

type User struct {
	ID    int
	Name  string
	Email string
}

func NewUser(name, email string) *User {
	return &User{Name: name, Email: email}
}
