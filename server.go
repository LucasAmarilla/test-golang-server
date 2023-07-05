package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", home)

	fmt.Println("USANDO DROGA NA PORTA 42069")
	log.Fatal(http.ListenAndServe(":42069", nil))
}
