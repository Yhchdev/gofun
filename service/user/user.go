package user

import . "gofun/model"

type Service interface {
	Register(name, password, email string) (string, error)
	Login(name, password string) (string, error)
}

type userService struct {
	user *User
}

func New() Service {
	return &userService{
		user: &User{},
	}
}

func (u *userService) Register(userName, userPassword, email string) (string, error) {

	return "", nil
}

func (*userService) Login(userName, userPassword string) (string, error) {
	return "", nil
}
