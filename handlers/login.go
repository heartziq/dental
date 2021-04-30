package handlers

import (
	"html/template"
	"log"
	"net/http"

	"dental/helper"

	uuid "github.com/satori/go.uuid"
)

type Login struct {
	Tpl *template.Template
}

func (l *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := ""
	username := r.URL.Query().Get("username")
	var userAlreadyExist string
	if r.URL.Query().Get("exist") == "true" {
		userAlreadyExist = "User already exist, pls login"
	}
	log.Println(username)

	if r.Method == http.MethodPost {
		// access conn from here

		username = r.FormValue("Username")

		user, err := helper.GetUser(conn, username)

		// if user not found, register first
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "/register?username="+username, http.StatusSeeOther)
			return
		}

		userInputPwd := r.FormValue("Password")

		// proceed with auth
		r.SetBasicAuth(username, userInputPwd)

		_, _, ok := r.BasicAuth()

		if ok {
			// auth successful
			// if p == user.Password {
			// 	// set cookie
			// 	id := uuid.NewV4()
			// 	userCookie := &http.Cookie{
			// 		Name:  "userInfo",
			// 		Value: id.String(),
			// 	}

			// 	http.SetCookie(w, userCookie)
			// 	// update session and redirect to /browse
			// 	helper.UpdateSession(user.UserName, userCookie.Value)
			// 	http.Redirect(w, r, "/browse", http.StatusSeeOther)
			// 	return
			// }

			isCorrectPwd := helper.VerifyPassword([]byte(user.Password), userInputPwd)

			if isCorrectPwd {
				id := uuid.NewV4()
				userCookie := &http.Cookie{
					Name:  "userInfo",
					Value: id.String(),
				}

				log.Printf("[login] Login passed: Isadmin: %v\n", user.IsAdmin)

				http.SetCookie(w, userCookie)
				// update session and redirect to /browse
				helper.UpdateSession(user.UserName, userCookie.Value)
				http.Redirect(w, r, "/browse", http.StatusSeeOther)
				return
			} else {
				log.Println("p not = user.Password")
				message = "Wrong Username/Password Combination"
			}

		}

	}

	log.Printf("message: %s\t[login.go, 72]\n", message)

	// GET method
	err := l.Tpl.ExecuteTemplate(w, "index.gohtml", struct {
		Error    string
		UserName string
		Exist    string
	}{
		message,
		username,
		userAlreadyExist,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

}
