package helper

import (
	"dental/handlers/data"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func GetUser(username string) (*data.User, error) {
	// fmt.Println(username, data.UserDB["tarar@aol.com"])
	user, exist := data.UserDB[username]

	if !exist {
		return nil, errors.New("user not found")
	}

	return &user, nil

}

func GetAllUser() (u []string) {
	for key, _ := range data.UserDB {
		if key == "admin" {
			continue
		}
		u = append(u, key)
	}

	return
}

func DisplayAllSession() {
	fmt.Println(data.SessionToUsername)
}

// to register new user
func AddUser(username string, pwd string) {
	data.UserDB[username] = data.User{
		Username: username,
		Password: pwd,
	}
	// log.Println(data.UserDB)
}

func SessionExist(session string) bool {
	_, exist := data.SessionToUsername[session]

	return exist
}

func IsLoggedIn(username string, session string) bool {

	if data.UsernameToSession[username] == session {
		return true
	}
	return false
}

func DeleteUser(username string) bool {
	_, exist := data.UserDB[username]
	if exist {
		delete(data.UserDB, username)
		return true
	}
	return exist
}

func UpdateSession(username, session string) {
	data.UsernameToSession[username] = session
	data.SessionToUsername[session] = username

	b, err := json.Marshal(data.SessionToUsername)
	if err != nil {
		fmt.Println("error:", err)
	}
	log.Println(string(b), "-----------> [helper.go, 60]")

	e, err := json.Marshal(data.UserDB)
	if err != nil {
		fmt.Println("error:", err)
	}

	log.Println(string(e), "-----------> [helper.go, 66]")
	ioutil.WriteFile("handlers/data/session.json", b, 0644)
	ioutil.WriteFile("handlers/data/user.json", e, 0644)

}

func Logout(session string) {
	username := data.SessionToUsername[session]
	delete(data.SessionToUsername, session)
	data.UsernameToSession[username] = ""
}

func GetNameBySession(session string) string {
	return data.SessionToUsername[session]
}
