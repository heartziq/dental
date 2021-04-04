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

	username := helper.GetNameBySession(cookie.Value)
	user, err := helper.GetUser(username)

	if err == nil {

		// force privilege escalation (debugging purpose)
		// user.IsAdmin = true
		var apptList []data.Appointment

		if !user.IsAdmin {
			// User load here
			user.Appointments = data.ListOne(user.Username)
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
