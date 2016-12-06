package main

import (
	"github.com/atakanozceviz/MVCTemplateTest/viewmodels"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	rtr := httprouter.New()
	rtr.GET("/", indexHandler)
	rtr.ServeFiles("/public/*filepath", http.Dir("public/"))
	log.Fatal(http.ListenAndServe(":80", rtr))
}

func indexHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	type pageData struct {
		Title string

	}

	clone, _ := templates.Clone()
	// you access the cached templates with the defined name, not the filename
	err := clone.ExecuteTemplate(w, "index", viewmodels.GetUsers())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
