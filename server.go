package main

import (
	"log"
	"net/http"
	"text/template"

	"dental/handlers"
)

// GET json path

var (
	login     *handlers.Login        = &handlers.Login{}
	logout    *handlers.Logout       = &handlers.Logout{}
	index     *handlers.Index        = &handlers.Index{}
	book      *handlers.Book         = &handlers.Book{}
	register  *handlers.Register     = &handlers.Register{}
	dashboard *handlers.ControlPanel = &handlers.ControlPanel{}
)

func init() {
	// set templates
	login.Tpl = template.Must(template.ParseGlob("templates/login/*"))
	index.Tpl = template.Must(template.ParseGlob("templates/browse/*"))
	book.Tpl = template.Must(template.ParseGlob("templates/book/*"))
	register.Tpl = template.Must(template.ParseGlob("templates/register/*"))
	dashboard.Tpl = template.Must(template.ParseGlob("templates/dashboard/*"))

}

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if no cookie, set cookie
		_, err := r.Cookie("userInfo")

		if r.URL.Path == "/login" || r.URL.Path == "/register" {
			if err != nil {
				log.Printf("Url: %v\tProceed\n", r.URL.Path)

				next.ServeHTTP(w, r)
				return
			}
			log.Printf("Url: %v\tRedir to /browse\n", r.URL.Path)

			http.Redirect(w, r, "/browse", http.StatusSeeOther)
		} else {
			if err != nil {
				log.Printf("Url: %v\tRedir to /login\n", r.URL.Path)

				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			log.Printf("Url: %v\tProceed\n", r.URL.Path)
			next.ServeHTTP(w, r)
		}
	})
}

func main() {

	customMux := http.NewServeMux()

	// Set Routers
	customMux.Handle("/", exampleMiddleware(index))
	customMux.Handle("/login", exampleMiddleware(login))
	customMux.Handle("/register", exampleMiddleware(register))
	customMux.Handle("/book", exampleMiddleware(book))
	customMux.Handle("/logout", exampleMiddleware(logout))
	customMux.Handle("/dashboard", exampleMiddleware(dashboard))

	customMux.Handle("/favicon.ico", http.NotFoundHandler())

	if err := http.ListenAndServe(":8080", customMux); err != nil {
		log.Fatal(err.Error())
	}

	// http.ListenAndServe(":8080", customMux)
}
