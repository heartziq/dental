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
	book  *handlers.Book  = &handlers.Book{}
)

func init() {
	// set templates
	login.Tpl = template.Must(template.ParseGlob("templates/login/*"))
	index.Tpl = template.Must(template.ParseGlob("templates/browse/*"))
	book.Tpl = template.Must(template.ParseGlob("templates/book/*"))

}

func main() {

	http.HandleFunc("/", index.HandleIndex)
	http.HandleFunc("/login", login.HandleLogin)
	http.HandleFunc("/book", book.HandleBook)

	// http.Handle("/", http.HandlerFunc(handlers.Login))

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
