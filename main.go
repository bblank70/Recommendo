package main

import (
	"net/http"
	"text/template"
)

var data = "This is rendering from Golang"
var tpl *template.Template

// init instantiates the templates, they must be .tmpl extenstions
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	// these are our paths
	http.HandleFunc("/", index)
	// http.HandleFunc("/verify", verifyer)
	// http.HandleFunc("/response", responder)
	// //this starts the server
	http.ListenAndServe(":8080", nil)

}

// ////////////////////////////////////////////

// index handles the route to the home page
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "base.html", data)
}
