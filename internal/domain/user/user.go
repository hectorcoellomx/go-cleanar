package user

// podría estar dentro de una carpeta model domain/user/model/user.go
type User struct {
	ID       int
	Username string
	Email    string
	Password string
}
