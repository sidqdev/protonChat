package storage

import (
	"github.com/google/uuid"
)

type User struct {
	UserName string `json:"username"`
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
		if u.UserName == user.UserName {
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

func (us *UserStorage) GetUserName(userId string) (string, bool) {
	for _, u := range us.Users {
		if u.UserID == userId {
			return u.UserName, true
		}
	}
	return "", false
}

func (us *UserStorage) CheckUserExist(username string) bool {
	for _, u := range us.Users {
		if u.UserName == username {
			return true
		}
	}
	return false
}
