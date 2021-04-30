package handlers

import (
	"database/sql"
	"dental/helper"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// load db
	conn   *sql.DB
	sqlErr error
)

func init() {
	conn, sqlErr = sql.Open("mysql", "user1:password@tcp(127.0.0.1:3306)/MYSTOREDB")
	if sqlErr != nil {
		log.Println(sqlErr.Error())
	}
}

type Register struct {
	Tpl *template.Template
}

func createPassword(password string) []byte {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost); err != nil {
		log.Println(err)
		return nil
	} else {
		return hash
	}
}

func (reg *Register) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var message string
	username := r.URL.Query().Get("username")
	if username != "" {
		message = "user not found, please register first"
	}
	if r.Method == http.MethodPost {

		// check if user exist
		username = r.FormValue("Username")
		_, err := helper.GetUser(conn, username)
		if err == nil {
			// user exist
			log.Println("user exist - redirect to /login")
			http.Redirect(w, r, "/login?username="+username+"&exist=true", http.StatusSeeOther)
			return
		}

		pwd := r.FormValue("Password")
		cPwd := r.FormValue("ConfirmPassword")

		if pwd == "" {
			message = "Password cannot be empty"
		} else {
			if pwd == cPwd {
				// cookie, _ := r.Cookie("userInfo")
				// helper.AddUser(username, pwd)
				// no cookie here - hence nil derefer
				helper.InsertRecord(conn, username, pwd, false)
				// set cookie
				id := uuid.NewV4()
				userCookie := &http.Cookie{
					Name:  "userInfo",
					Value: id.String(),
				}

				http.SetCookie(w, userCookie)
				helper.UpdateSession(username, userCookie.Value)
				http.Redirect(w, r, "/browse", http.StatusSeeOther)
			} else {
				message = "Password and Confirm Password do not match"
			}
		}

	}

	err := reg.Tpl.ExecuteTemplate(w, "index.gohtml", struct {
		Error    string
		UserName string
	}{
		message,
		username,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
