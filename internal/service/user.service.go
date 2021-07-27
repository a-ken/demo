package service

import "errors"

type UserService struct {
}

func (u *UserService) Signin(name string, password string) (bool, error) {
	if name != "test" {
		return false, errors.New("user not found")
	}
	if password != "abcd1234" {
		return false, errors.New("incorrect password")
	}
	return true, nil
}