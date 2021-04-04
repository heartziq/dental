package helper

import (
	"dental/handlers/data"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

// GetUser retrieve a user based on username provided as param
// if user does not exist,
// error "user not found" will be return alongside a nil obj
//
func GetUser(username string) (*data.User, error) {
	user, exist := data.UserDB[username]

	if !exist {
		return nil, errors.New("user not found")
	}

	return &user, nil

}

// GetAllUser will retrieve all registered user as a slice (except admin)
//
func GetAllUser() (u []string) {
	for key := range data.UserDB {
		if key == "admin" {
			continue
		}
		u = append(u, key)
	}

	return
}

// AddUser adds new user account to UserDB
// The 2nd arg 'pwd' will be the user's password
func AddUser(username string, pwd string) {
	data.UserDB[username] = data.User{
		Username: username,
		Password: pwd,
	}
}

// SessionExist verify if session is in map
// Returns true if exist, false otherwise
func SessionExist(session string) bool {
	_, exist := data.SessionToUsername[session]

	return exist
}

func IsLoggedIn(username string, session string) bool {

	return data.UsernameToSession[username] == session
}

func DeleteUser(username string) bool {
	_, exist := data.UserDB[username]
	if exist {
		delete(data.UserDB, username)
		e, err := json.Marshal(data.UserDB)
		if err != nil {
			fmt.Println("error:", err)
		}
		ioutil.WriteFile("handlers/data/user.json", e, 0644)
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

	// set to empty session
	data.UsernameToSession[username] = ""

	// Update
	b, err := json.Marshal(data.SessionToUsername)
	if err != nil {
		fmt.Println("error:", err)
	}

	log.Println(string(b), "-----------> [helper.go, 103]")
	ioutil.WriteFile("handlers/data/session.json", b, 0644)

}

func GetNameBySession(session string) string {
	return data.SessionToUsername[session]
}
