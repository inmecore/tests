package model

import (
	"errors"
	"tests/lib"
	"time"
)

var UserModel = User{}

type User struct {
	Id          int        `json:"id"`
	Username    string     `json:"username"`
	Name        string     `json:"name"`
	Phone       string     `json:"phone"`
	Avatar      string     `json:"avatar"`
	Password    string     `json:"password"`
	Description string     `json:"description"`
	LastLoginAt *time.Time `json:"lastLoginAt"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

var UserInfo = User{
	Id:          1,
	Username:    "test",
	Name:        "Test",
	Phone:       "13123456789",
	Avatar:      "https://www.xxx.com/xxx.jpg",
	Password:    "FF1757E51EC5E7F5D719B4623A4946BB",
	Description: "test user info",
	LastLoginAt: &lib.Now,
	CreatedAt:   &lib.Now,
	UpdatedAt:   &lib.Now,
	DeletedAt:   nil,
}

func (*User) Login(username, password string) (bool, error) {
	return UserInfo.Username == username && UserInfo.Password == lib.Password(password), nil
}

func (*User) Info(uid int) (*User, error) {
	if uid != UserInfo.Id {
		return nil, errors.New("not found")
	}
	return &UserInfo, nil
}
