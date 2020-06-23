package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./client/*.html"))
}

func main() {
	fs := http.FileServer(http.Dir("./client"))
	http.Handle("/", fs)
	http.HandleFunc("/about", about)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("localhost:8080", nil)
}

func serveClient() {
	fs := http.FileServer(http.Dir("./client"))
	http.Handle("/", fs)
}

func about(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)
}
