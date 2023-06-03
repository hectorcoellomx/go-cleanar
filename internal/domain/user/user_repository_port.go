package user

type UserRepositoryPort interface {
	FindAll() (*[]User, error)
	Create(user *User) error
	FindByID(id uint) (*User, error)
}
