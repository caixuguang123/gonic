package main

import (
	"qiniupkg.com/x/errors.v7"

)

type user struct {
	ID     int `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

var id int
func registerUser(userName string, password string) (*user, error) {
	if isUserNameAvailable(userName) {
		return nil, errors.New("user dumplicate")
	}
	if password == "" {
		return nil, errors.New("pwd can't be empty")
	}
	*id = id+1;
	user := user{id, userName, password}
	store(&user)
	return &user, nil
}

func isUserNameAvailable(userName string) bool {
	return GetbyName(userName)
}

func isUserValid(userName, password string) bool {
	return GetbyUser(userName,password)
}