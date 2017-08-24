package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine"

	"SimplystDream/simplyst"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	signupTmpl = parseTemplate("signup.html")
)

func main() {

	registerHandlers()
	appengine.Main()
}

func registerHandlers() {

	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/signup", http.StatusFound))
	r.Methods("GET").Path("/signup").
		Handler(appHandler(addFormHandler))
	r.Methods("POST").Path("/signup").Handler(appHandler(signupHandler))

	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}
func userFromForm(r *http.Request) (*simplyst.User, error) {

	user := &simplyst.User{
		FirstName: r.FormValue("fname"),
		LastName:  r.FormValue("lname"),
		Email:     r.FormValue("email"),
		UserName:  r.FormValue("username"),
	}
	return user, nil
}

func signupHandler(w http.ResponseWriter, r *http.Request) *appError {
	user, err := userFromForm(r)
	if err != nil {
		return appErrorf(err, "could not parse user from form: %v", err)
	}
	id, err := simplyst.DB.Adduser(user)
	if err != nil {
		return appErrorf(err, "could not save user: %v", err)
	}
	http.Redirect(w, r, fmt.Sprintf("/signup/%d", id), http.StatusFound)
	return nil
}

func addFormHandler(w http.ResponseWriter, r *http.Request) *appError {
	return signupTmpl.Execute(w, r, nil)
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)

		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}
