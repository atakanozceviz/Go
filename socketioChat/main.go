package main

import (
	"database/sql"
	"log"
	"net/http"

	"encoding/json"
	"github.com/googollee/go-socket.io"
	_ "github.com/mattn/go-sqlite3"
)

var db, _ = sql.Open("sqlite3", "users.db")
var onlineUsers = make(map[string]string)

func main() {

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {

		so.On("subscribe", func(room string) {
			so.Join(room)
			onlineUsers[so.Id()] = room
			so.BroadcastTo("admin", "ToAdmin", "refresh")
		})

		so.On("send message", func(data map[string]string) {
			msg := data["message"]
			room := data["room"]
			if room != "admin" {
				so.BroadcastTo("admin", "ToAdmin", data)
			}

			so.BroadcastTo(room, "conversation private post", msg)
		})

		so.On("change room", func(data map[string]string) {
			newroom := data["newroom"]
			oldroom := so.Rooms()
			if len(oldroom) > 0 {
				for i := range oldroom {
					if oldroom[i] != "admin" {
						log.Println("Leaving room: ", oldroom[i])
						so.Leave(oldroom[i])
					}

				}
			}
			so.Join(newroom)
			log.Println(so.Rooms())
		})

		so.On("disconnection", func() {
			delete(onlineUsers, so.Id())
			so.BroadcastTo("admin", "ToAdmin", "refresh")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	http.HandleFunc("/admin", serveStatic)
	http.HandleFunc("/lookup", lookup)
	http.HandleFunc("/whoisOnline", whoisOnline)
	log.Println("Serving at localhost:80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "asset/admin.html")
}

func lookup(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	rows, err := db.Query("SELECT id FROM user WHERE email = ? ", email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var id string

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	defer rows.Close()

	w.Write([]byte(id))
}
func whoisOnline(w http.ResponseWriter, r *http.Request) {

	type idEmail struct {
		ID    string `json:"id,attr"`
		Email string `json:"email,attr"`
	}
	got := []idEmail{}

	for _, id := range onlineUsers {
		rows, err := db.Query("SELECT * FROM user WHERE id = ? ", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}

		for rows.Next() {
			var p idEmail
			err = rows.Scan(&p.ID, &p.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println(err)
				return
			}
			got = append(got, p)
		}
		rows.Close()
	}

	pJ, err := json.Marshal(got)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(string(pJ))
}
