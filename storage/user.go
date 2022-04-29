package storage

import (
	"github.com/google/uuid"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserID   string `json:"UserID"`
}

type UserStorage struct {
	Users []User `json:"users"`
}

func (us *UserStorage) LoginUser(user User) (bool, string) {
	uuid4, _ := uuid.NewUUID()
	userId := uuid4.String()

	for _, u := range us.Users {
		if u.Login == user.Login {
			if u.Password == user.Password {
				u.UserID = userId
				return true, userId
			} else {
				return false, ""
			}
		}
	}

	user.UserID = userId
	us.Users = append(us.Users, user)
	return true, userId
}

func (us *UserStorage) GetLogin(userId string) (bool, string) {
	for _, u := range us.Users {
		if u.UserID == userId {
			return true, u.Login
		}
	}
	return false, ""
}
