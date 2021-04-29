package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"dental/handlers/data"
	"dental/helper"

	uuid "github.com/satori/go.uuid"
)

const (
	dateTimeFmt = "2006-01-02 15:04"
	altFmt      = "Jan 02, 2006 at 3:00pm (SGT)"
)

var (
	newDate date
)

type Book struct {
	Tpl *template.Template
}

type date struct {
	Start string
	End   string
	Appt  string
	Time  string
}

func (b *Book) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("userInfo")
	username := helper.GetNameBySession(cookie.Value)

	// Admin will be redirected to dashboard
	if username == "admin" {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	user, err := helper.GetUser(username)
	log.Printf("[/book] username: %s\tuser: %v\n", username, user)
	if err == nil {
		start := time.Now().Format("2006-01-02")                         // yyyy-dd-mm
		end := time.Now().Add(time.Hour * 24 * 182).Format("2006-01-02") // approx 6 months.
		newDate = date{
			Start: start,
			End:   end,
		}

		if r.Method == http.MethodPost {

			// update date value
			newDate.Appt = r.FormValue("appt-date")
			newDate.Time = r.FormValue("appt-time")

			usrSelected := newDate.Appt + " " + newDate.Time
			t, err := time.Parse(dateTimeFmt, usrSelected)
			if err != nil {
				fmt.Println(err.Error())
			}

			newDate.Appt = t.Format(altFmt)

			newAppt := data.Appointment{
				Id:       uuid.NewV4().String(),
				Customer: user.Username,
				Doctor:   "Kestart D",
				Time:     newDate.Appt,
				Location: "Wing A, Level 3",
			}

			data.B.Insert(&data.RootNode, &newAppt, user.Username)

			err = b.Tpl.ExecuteTemplate(w, "redir.gohtml", struct {
				Seconds int
				Message string
			}{
				2,
				"Booking Successful",
			})

			if err != nil {
				log.Fatal(err.Error())
			}

		}
	}

	err = b.Tpl.ExecuteTemplate(w, "index.gohtml", struct {
		Start    string
		End      string
		Appt     string
		Username string
		IsAdmin  bool
	}{
		newDate.Start,
		newDate.End,
		newDate.Appt,
		user.Username,
		user.IsAdmin,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
