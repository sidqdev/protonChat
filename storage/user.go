package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/google/uuid"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	UserID   string `json:"UserID"`
}

type UserStorage struct {
	Users []User `json:"users"`
	mu    sync.Mutex
}

func (us *UserStorage) Save() {
	us.mu.Lock()
	content, err := json.Marshal(us)
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile("UserStorage.json", content, 0644)
	if err != nil {
		log.Println(err)
	}
	us.mu.Unlock()
}

func (us *UserStorage) Load() {
	content, err := ioutil.ReadFile("UserStorage.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(content, &us)
	if err != nil {
		log.Println(err)
	}
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
	go us.Save()
	return true, userId
}

func (us *UserStorage) Logout(userId string) {
	for _, u := range us.Users {
		if u.UserID == userId {
			u.UserID = ""
		}
	}
	go us.Save()
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
