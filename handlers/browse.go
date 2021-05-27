package handlers

import (
	"html/template"
	"log"
	"net/http"

	"dental/handlers/data"
	"dental/helper"
)

type Index struct {
	Tpl *template.Template
}

// Handle Index provide info on already booked appointments
// Customer will see ONLY THEIR list of booked appointments
// Admin will see the whole entire list (from ALL customers)
func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("userInfo")

	username := helper.GetNameBySession(cookie.Value)
	log.Printf(".GetNameBySession -> %v\n", username)
	user, err := helper.GetUser(conn, username)
	log.Printf(".GetUser -> %v\n", user)

	if err == nil {

		// force privilege escalation (debugging purpose)
		// user.IsAdmin = true
		var apptList []data.Appointment

		if !user.IsAdmin {
			// User load here
			user.Appointments = data.ListOne(user.UserName)
		} else {

			// display all appointments
			data.B.Display(data.RootNode, &apptList)
			user.Appointments = apptList
		}

		// only Admin can search
		if r.Method == http.MethodPost {

			newName := r.FormValue("search")
			if newName != "" {
				d := &data.Appointment{}
				data.B.Search(&data.RootNode, newName, &d)
				if d.Customer != "" {
					user.Appointments = []data.Appointment{*d}

				}
			}

		}
	}

	log.Printf("[/browse] username: %s\tuser: %v\n", username, user)

	err = i.Tpl.ExecuteTemplate(w, "index.gohtml", user)
	if err != nil {
		log.Fatal(err.Error())
	}
}
