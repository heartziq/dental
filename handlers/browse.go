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

// Handle Index provide info of already booked appointment
// Customer will see their list of booked appointment
// Admin will see the whole entire list (from all customers)
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

		// force privilege escalation
		user.IsAdmin = true

		if !user.IsAdmin {
			// User load here
			user.Appointments = data.ListOne(user.Username)
		}

		// only Admin can search
		if r.Method == http.MethodPost {

			newName := r.FormValue("search")
			d := &data.Appointment{}
			data.B.Search(&data.RootNode, newName, &d)

			user.Appointments = []data.Appointment{*d}

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
