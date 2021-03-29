package main

import (
	"net/http"
	"text/template"

	"dental/handlers"
	_ "dental/handlers/data"
)

var (
	login *handlers.Login = &handlers.Login{}
	index *handlers.Index = &handlers.Index{}
)

func init() {
	// set templates
	login.Tpl = template.Must(template.ParseGlob("templates/login/*"))
	index.Tpl = template.Must(template.ParseGlob("templates/browse/*"))

}

func main() {

	http.HandleFunc("/", index.HandleIndex)
	http.HandleFunc("/login", login.HandleLogin)

	// http.Handle("/", http.HandlerFunc(handlers.Login))

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
