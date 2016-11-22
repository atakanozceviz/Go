package main

import (
	"net/http"

	"github.com/atakanozceviz/Gophr/templates"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.RenderTemplate(w, r, "index/home", nil)
	})

	mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		templates.RenderTemplate(w, r, "layout", nil)
	})

	mux.Handle(
		"/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))),
	)

	http.ListenAndServe(":80", mux)
}
