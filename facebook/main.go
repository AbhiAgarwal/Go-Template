package main

import (
	"net/http"

	. "github.com/abhiagarwal/go-template/facebook/handlers"
	negroni "github.com/codegangsta/negroni"
	oauth2 "github.com/goincremental/negroni-oauth2"
	sessions "github.com/goincremental/negroni-sessions"
	cookiestore "github.com/goincremental/negroni-sessions/cookiestore"
	mux "github.com/gorilla/mux"
)

const (
	Address         string = ":3000"
	StaticDirectory string = "./public/"
)

// Main
func main() {
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir(StaticDirectory)))

	store := cookiestore.New([]byte("secret123"))
	n.Use(sessions.Sessions("my_session", store))
	n.Use(oauth2.Facebook(&oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "",
		Scopes:       []string{},
	}))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", MainPageHandler)

	n.UseHandler(router)
	n.Run(Address)
}
