package handlers

import (
	"dental/helper"
	"html/template"
	"net/http"
)

type Logout struct {
	Tpl *template.Template
}

func (logout *Logout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("userInfo")

	cookie.MaxAge = -1
	helper.Logout(cookie.Value)
	http.SetCookie(w, cookie) // cookie destroy by next request

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
