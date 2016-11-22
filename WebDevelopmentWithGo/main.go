package main

import (
	"html/template"
	"log"
	"net/http"

	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))

type TemplateData struct {
	Title  string
	Active string
}


func main() {

	db, _ := sql.Open("sqlite3", "dev.db")
	fmt.Println(db.Ping())

	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.GET("/register", RegisterHandler)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/404.html")
	})
	router.ServeFiles("/public/*filepath", http.Dir("public/"))
	log.Fatal(http.ListenAndServe(":80", router))
}

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data := TemplateData{
		Title: "Home Page",
	}

	clone, _ := templates.Clone()

	// you access the cached templates with the defined name, not the filename
	err := clone.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data := TemplateData{
		Title:  "Register",
		Active: "register",
	}

	clone, _ := templates.Clone()

	// you access the cached templates with the defined name, not the filename
	err := clone.ExecuteTemplate(w, "register", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


