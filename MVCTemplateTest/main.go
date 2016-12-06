package main

import (
	"github.com/atakanozceviz/MVCTemplateTest/viewmodels"
	"github.com/julienschmidt/httprouter"
	"github.com/leekchan/gtf"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

// compile all templates and cache them
var templates = template.Must(template.New("").Funcs(gtf.GtfFuncMap).ParseGlob("templates/*"))

func main() {
	rtr := httprouter.New()
	rtr.GET("/", indexHandler)
	rtr.ServeFiles("/public/*filepath", http.Dir("public/"))
	log.Fatal(http.ListenAndServe(":80", rtr))
}

func indexHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	type pageData struct {
		Title string
		Users viewmodels.Users
	}
	pd := pageData{
		Title: "Show Generated Data",
		Users: viewmodels.GetUsers(),
	}
	clone, _ := templates.Clone()
	// you access the cached templates with the defined name, not the filename
	err := clone.ExecuteTemplate(w, "index", pd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
