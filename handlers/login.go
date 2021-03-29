package handlers

import (
	"log"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

type Login struct {
	Tpl *template.Template
}

type User struct {
	Username string
	Password string
}

var (
	user     User = User{}
	mapUsers      = map[string]User{} // key is session
)

func (l *Login) HandleLogin(w http.ResponseWriter, r *http.Request) {

	// if cookie exist, redir to /
	_, err := r.Cookie("userInfo")
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		// set cookie
		id := uuid.NewV4()
		userCookie := &http.Cookie{
			Name:  "userInfo",
			Value: id.String(),
		}

		http.SetCookie(w, userCookie)

		// insert new user
		user.Username = r.FormValue("Username")
		user.Password = r.FormValue("Password")

		mapUsers[userCookie.Value] = user

		// load /
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

	// GET method with No Cookie? Load empty user{}
	err = l.Tpl.ExecuteTemplate(w, "index.gohtml", user)
	if err != nil {
		log.Fatal(err.Error())
	}
}