package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
	Desc  string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{
		"Lucas",
		"é canibal",
		"Arrebeta sua bunda no s3x0 4n4L",
	}

	templates.ExecuteTemplate(w, "index.html", u)
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", home)

	fmt.Println("USANDO DROGA NA PORTA 42069")
	log.Fatal(http.ListenAndServe(":42069", nil))
}
