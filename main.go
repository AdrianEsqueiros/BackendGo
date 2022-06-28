package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/create", Create)
	log.Println("Starting server")
	http.ListenAndServe(":8080", nil)
}

func Start(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome")
	templates.ExecuteTemplate(w, "start", nil)
}
func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}
