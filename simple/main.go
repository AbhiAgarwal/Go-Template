package main

import (
	"net/http"

	. "github.com/abhiagarwal/go-template/simple/handlers"
	negroni "github.com/codegangsta/negroni"
)

const (
	Address         string = ":3000"
	StaticDirectory string = "./public/"
)

// Main
func main() {
	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir(StaticDirectory)))

	router := http.NewServeMux()
	router.HandleFunc("/", MainPageHandler)

	n.UseHandler(router)
	n.Run(Address)
}
