package data

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	UserName, Password string
	IsAdmin            bool
	Appointments       []Appointment
}

var (
	UserDB            = map[string]User{}
	UsernameToSession = map[string]string{} //key is username, value = session id
	SessionToUsername = map[string]string{}
)

func init() {
	// pre-load previously stored user session
	content, err := os.ReadFile("handlers/data/session.json")
	if err != nil {
		log.Println(err.Error())
	}

	json.Unmarshal(content, &SessionToUsername)

	// pre-load previously stored user db
	content, err = os.ReadFile("handlers/data/user.json")
	if err != nil {
		log.Println(err.Error())
	}

	json.Unmarshal(content, &UserDB)
}
