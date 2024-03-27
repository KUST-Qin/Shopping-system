package user

import (
	"errors"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(30)"`
	Password  string `gorm:"type:varchar(100)"`
	Password2 string `gorm:"-"`
	Salt      string `gorm:"type:varchar(100)"`
	Token     string `gorm:"type:varchar(500)"`
	IsDeleted bool
	IsAdmin   bool
}

// 新建用户实例
func NewUser(username, password, password2 string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Password2: password2,
		IsDeleted: false,
		IsAdmin:   false,
	}
}

// 用户自定义错误
var (
	ErrUserExistWithName   = errors.New("用户名已经存在")
	ErrUserNotFound        = errors.New("用户未找到")
	ErrMismatchedPasswords = errors.New("密码不匹配")
	ErrInvalidUsername     = errors.New("无效用户名")
	ErrInvalidPassword     = errors.New("无效密码")
)
