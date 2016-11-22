package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func rootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}
}

func main() {
	var host = flag.String("host", "127.0.0.1", "IP of host to run webserver on")
	var port = flag.Int("port", 80, "Port to run webserver on")
	var staticPath = flag.String("staticPath", "static/", "Path to static files")

	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(*staticPath))))

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening on %s", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
