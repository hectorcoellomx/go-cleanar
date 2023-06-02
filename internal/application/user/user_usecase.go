package user

type UserUseCase interface {
	CreateUser(name string, email string) (*User, error)
	GetUserByID(id uint) (*User, error)
}

type userUseCase struct {
	userService UserService
}

func NewUserUseCase(userService UserService) UserUseCase {
	return &userUseCase{
		userService: userService,
	}
}

func (uc *userUseCase) CreateUser(name string, email string) (*User, error) {
	return uc.userService.CreateUser(name, email)
}

func (uc *userUseCase) GetUserByID(id uint) (*User, error) {
	return uc.userService.GetUserByID(id)
}
