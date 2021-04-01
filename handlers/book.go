package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"dental/handlers/data"

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

func (b *Book) HandleBook(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("userInfo")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	currentUser, exist := mapUsers[cookie.Value]
	if exist {
		// fmt.Println(currentUser)
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

			// (data.RootNode, &user.Appointments)

			newAppt := data.Appointment{
				Id:       uuid.NewV4().String(),
				Customer: currentUser.Username,
				Doctor:   "Kestart D",
				Time:     newDate.Appt,
				Location: "Wing A, Level 3",
			}

			data.B.Insert(&data.RootNode, &newAppt, currentUser.Username)
			fmt.Println("Book.go")
			data.B.Display(data.RootNode)

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
	} else {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	fmt.Println("After block: ", currentUser.Username)

	err = b.Tpl.ExecuteTemplate(w, "index.gohtml", struct {
		Start    string
		End      string
		Appt     string
		Username string
	}{
		newDate.Start,
		newDate.End,
		newDate.Appt,
		currentUser.Username,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
