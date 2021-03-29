package handlers

import (
	"dental/handlers/data"
	"log"
	"net/http"
	"text/template"
)

type Index struct {
	Tpl *template.Template
}

func (i *Index) HandleIndex(w http.ResponseWriter, r *http.Request) {
	// if not logged in, redir to "/login"
	cookie, err := r.Cookie("userInfo")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// retrieve user
	if currentUser, exist := mapUsers[cookie.Value]; exist {
		user = currentUser

		// load appointments here
		user.Appointments = []data.Appointment{}
		data.B.List(data.RootNode, &user.Appointments)
		if r.Method == http.MethodPost {
			newName := r.FormValue("search")
			user.Username = newName
			mapUsers[cookie.Value] = user
		}
	} else {
		// cookie exist but user not found
		// reset cookie to prevent infinite redirect
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	err = i.Tpl.ExecuteTemplate(w, "index.gohtml", user)
	if err != nil {
		log.Fatal(err.Error())
	}
}
