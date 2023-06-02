package user

// Podría estar en la carpeta ports domain/user/ports/out/repository.go
// podría llamarse UserRepositoryPort

type Repository interface {
	Create(user *User) error
	FindByID(id uint) (*User, error)
	// Otros métodos del repositorio necesarios para las operaciones de usuario
}
