package main

import (
	"net/http"
	"os"

	"fmt"

	"github.com/gorilla/mux"
)

const (
	Port = ":80"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	rtr.HandleFunc("/{id:.+}", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(Port, nil)

}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	_, err := os.Stat(fileName)
	if err != nil {
		fileName = "files/404.html"
	}

	fmt.Println(pageID, fileName)

	http.ServeFile(w, r, fileName)
}
