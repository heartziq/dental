package handlers

import (
	"html/template"
	"log"
	"net/http"

	"dental/helper"
)

type ControlPanel struct {
	Tpl *template.Template
}

func (cp *ControlPanel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("userInfo")
	currUser := helper.GetNameBySession(cookie.Value)
	if currUser != "admin" {
		log.Println("Only admin can edit user")
		http.Redirect(w, r, "/browse", http.StatusSeeOther)
		return
	}
	username := r.URL.Query().Get("username")
	// var userList = []{}
	helper.DeleteUser(conn, username)
	// list all
	userList := helper.GetAllUser(conn)

	err := cp.Tpl.ExecuteTemplate(w, "dashboard.gohtml", struct {
		UserList []string
	}{
		userList,
	})

	if err != nil {
		log.Println("Error loading template: templates/dashboard/dashboard.gohtml")
	}
}
