package viewmodels

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID      int
	Name    string
	Address string
	Age     int
	ImgURL  string
}

type Users []User

var db, _ = sql.Open("sqlite3", "users.db")

func GetUsers() Users {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
	}

	var users Users
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Address, &u.Age, &u.ImgURL)
		if err != nil {
			log.Println(err)
		}
		users = append(users, u)
	}
	defer rows.Close()
	return users
}
