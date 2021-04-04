package handlers

import (
	"log"
	"net/http"
	"text/template"

	"dental/handlers/data"
	"dental/helper"
)

type Index struct {
	Tpl *template.Template
}

// Handle Index provide info of already booked appointment
// Customer will see their list of booked appointment
// Admin will see the whole entire list (from all customers)
func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("userInfo")
	helper.DisplayAllSession()

	// return "" if new session
	username := helper.GetNameBySession(cookie.Value)
	log.Printf("usrename: %s\t[browse.go, 22]\n", username)
	// retrieve user (now user is not persisted)
	user, err := helper.GetUser(username)
	log.Printf("user: %v\t[browse.go, 28]\n", user)

	if err == nil {

		// force privilege escalation
		// user.IsAdmin = true

		if !user.IsAdmin {
			// User load here
			user.Appointments = data.ListOne(user.Username)
		}

		// only Admin can search
		if r.Method == http.MethodPost {

			newName := r.FormValue("search")
			d := &data.Appointment{}
			data.B.Search(&data.RootNode, newName, &d)
			log.Printf("d: %v\t[browse.go, 46]\n", (*&d).Customer)
			if (*&d).Customer != "" {
				user.Appointments = []data.Appointment{*d}

			}
			log.Printf("len(appt): %v\t[browse.go, 48]\n", len(user.Appointments))

		}
	} else {
		// cookie exist but user not found
		// reset cookie to prevent infinite redirect
		cookie.MaxAge = -1
		http.SetCookie(w, cookie) // cookie destroy by next request
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	err = i.Tpl.ExecuteTemplate(w, "index.gohtml", user)
	if err != nil {
		log.Fatal(err.Error())
	}
}
