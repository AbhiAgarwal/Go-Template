package handler

import (
	"html/template"
	"net/http"
)

const (
	templates       string = "./templates/"
	partialTemplate string = templates + "partials/"
	baseTemplate    string = partialTemplate + "base.html"
	homeTemplate    string = templates + "home.html"
)

func MainPageHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	var logInTemplate = template.Must(template.New("home").ParseFiles(baseTemplate, homeTemplate))
	logInTemplate.ExecuteTemplate(res, "base", nil)
	return
}
