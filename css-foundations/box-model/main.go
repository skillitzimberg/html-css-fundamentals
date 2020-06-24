package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./site/*html"))
}

func main() {
	fs := http.FileServer(http.Dir("./site"))
	http.Handle("/", fs)
	http.HandleFunc("/border", border)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func border(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "border.html", nil)
}
