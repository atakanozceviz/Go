package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"

	"database/sql"
	"fmt"

	"encoding/json"

	"io/ioutil"

	"encoding/xml"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))
var db, _ = sql.Open("sqlite3", "dev.db")

type templateData struct {
	Title  string
	Active string
}

type searchResult struct {
	Title  string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Year   string `xml:"hyr,attr"`
	ID     string `xml:"owi,attr"`
}

type classifySearchResponse struct {
	Results []searchResult `xml:"works>work"`
}

type classifyBookResponse struct {
	BookData struct {
		Title  string `xml:"title,attr"`
		Author string `xml:"author,attr"`
		ID     string `xml:"owi,attr"`
	} `xml:"work"`
	Classification struct {
		MostPopular string `xml:"sfa,attr"`
	} `xml:"recommendations>ddc>mostPopular"`
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/register", registerHandler)
	router.POST("/search", searchHandler)
	router.GET("/books/add", booksAddHandler)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/404.html")
	})
	router.ServeFiles("/public/*filepath", http.Dir("public/"))
	log.Fatal(http.ListenAndServe(":80", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data := templateData{
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

func registerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data := templateData{
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

func searchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	results, err := search(r.FormValue("search"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func booksAddHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	book, err := find(r.FormValue("id"))

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err = db.Ping(); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = db.Exec("INSERT INTO books(pk, title, author, id, classification) VALUES (?, ?, ?, ?, ?)",
		nil, book.BookData.Title, book.BookData.Author, book.BookData.ID,
		book.Classification.MostPopular)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func search(query string) ([]searchResult, error) {

	body, err := classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&title=" + url.QueryEscape(query))
	if err != nil {
		return []searchResult{}, err
	}

	var c classifySearchResponse
	err = xml.Unmarshal(body, &c)
	return c.Results, err

}

func find(id string) (classifyBookResponse, error) {
	var c classifyBookResponse
	body, err := classifyAPI("http://classify.oclc.org/classify2/Classify?&summary=true&owi=" + url.QueryEscape(id))
	if err != nil {
		return classifyBookResponse{}, err
	}

	err = xml.Unmarshal(body, &c)
	return c, err

}

func classifyAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
