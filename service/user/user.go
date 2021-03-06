package user

import (
	"errors"
	"gofun/database/mysql"
	. "gofun/model"
	"gorm.io/gorm"
)

type Error struct {
	error
}

func userError(err string) Error {
	return Error{errors.New(err)}
}

var (
	Secret        = []byte("@fc6951544^f55c644!@0d")
	Existed       = userError("邮箱已存在")
	NotExisted    = userError("邮箱不存在")
	PasswordWrong = userError("邮箱或密码错误")
	LoginFailed   = userError("登录失败")
)

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

func (u *userService) Register(name, password, email string) (string, error) {

	user := User{
		Model:    gorm.Model{},
		Name:     name,
		Password: password,
		Email:    email,
	}

	result := mysql.DBConn().Create(user)
	if result.Error != nil {
		return "", result.Error
	}
	return "token", nil
}

func (*userService) Login(userName, userPassword string) (string, error) {
	return "", nil
}
