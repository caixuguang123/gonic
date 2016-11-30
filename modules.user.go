package main

import "errors"

type user struct {
	ID       int `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

var userList = []user{
	{1, "u1", "p1"},
	{2, "u2", "p2"},
	{3, "u3", "p3"},
}

func registerUser(userName string, password string) (*user, error) {
	if !isUserNameAvailable(userName) {
		return nil, errors.New("user dumplicate")
	}
	if password == "" {
		return nil, errors.New("pwd can't be empty")
	}
	id := len(userList) - 1;
	var user = user{id, userName, password}
	userList = append(userList, user)
	return &user, nil
}

func isUserNameAvailable(userName string) bool {
	for _, name := range userList {
		if name.UserName == userName {
			return false
		}
	}
	return true
}

func isUserValid(userName, password string) bool {
	for _, name := range userList {
		if name.UserName == userName {
			if name.Password == password {
				return true
			}
		}
	}
	return false
}