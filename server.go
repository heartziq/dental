package main

import (
	"html/template"
	"log"
	"net/http"

	"dental/handlers"
)

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

func VerifyLoggedIn(next http.Handler) http.Handler {
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
	customMux.Handle("/", VerifyLoggedIn(index))
	customMux.Handle("/login", VerifyLoggedIn(login))
	customMux.Handle("/register", VerifyLoggedIn(register))
	customMux.Handle("/book", VerifyLoggedIn(book))
	customMux.Handle("/logout", VerifyLoggedIn(logout))
	customMux.Handle("/dashboard", VerifyLoggedIn(dashboard))

	customMux.Handle("/favicon.ico", http.NotFoundHandler())

	if err := http.ListenAndServeTLS(":5221", "cert/cert.pem", "cert/key.pem", customMux); err != nil {
		log.Fatal(err.Error())
	}

}
