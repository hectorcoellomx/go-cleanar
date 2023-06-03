package user

// Podría estar en la carpeta ports domain/user/ports/out/repository.go
// podría llamarse UserRepositoryPort

type Repository interface {
	Get() (*[]User, error)
	Create(user *User) error
	FindByID(id uint) (*User, error)
}
