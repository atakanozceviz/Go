package main

import (
	"html/template"
	"log"
	"net/http"

	"database/sql"
	"fmt"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))

var db, _ = sql.Open("sqlite3", "db.db")

func main() {

	fmt.Println(db.Ping())

	router := httprouter.New()
	router.GET("/", indexHandler)
	router.POST("/table/:name", tableActions)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/404.html")
	})
	router.ServeFiles("/public/*filepath", http.Dir("public/"))
	log.Fatal(http.ListenAndServe(":80", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	clone, _ := templates.Clone()

	// you access the cached templates with the defined name, not the filename
	err := clone.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

type person struct {
	ID   int    `json:"id,attr"`
	Name string `json:"name,attr"`
	Age  int    `json:"age,attr"`
}

func tableActions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if ps.ByName("name") == "list" {

		rows, err := db.Query("SELECT * FROM person")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		got := []person{}
		for rows.Next() {
			var p person
			err = rows.Scan(&p.ID, &p.Name, &p.Age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err)
				return
			}
			got = append(got, p)
		}
		defer rows.Close()
		pJ, err := json.Marshal(got)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
		json.NewEncoder(w).Encode(string(pJ))

	} else if ps.ByName("name") == "edit" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		age := r.FormValue("age")
		fmt.Println(id, name, age)
		_, err := db.Exec("UPDATE person SET name = ?, age = ? WHERE id = ?", name, age, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
}
