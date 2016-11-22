package main

import "net/http"
import "time"
import "fmt"

const (
	Port = ":80"
)

func main() {
	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/", serveDynamic)
	http.ListenAndServe(Port, nil)
}

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "Şuan saat" + time.Now().String()
	fmt.Fprintln(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}
