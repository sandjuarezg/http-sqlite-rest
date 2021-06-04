package user

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type users struct {
	UsersArray []user `json:"users"`
}

func AddUser(db *sql.DB, body []byte) (err error) {
	var users users
	smt, err := db.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
	if err != nil {
		return
	}
	defer smt.Close()

	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Fatal(err)
	}

	_, err = smt.Exec(users.UsersArray[0].Name, users.UsersArray[0].Pass)
	if err != nil {
		return
	}

	return
}

func ShowUser(db *sql.DB) (users users, err error) {
	rows, err := db.Query("SELECT id, name, password FROM users")
	if err != nil {
		return
	}
	defer rows.Close()

	var user user

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Pass)
		if err != nil {
			return
		}
		users.UsersArray = append(users.UsersArray, user)
	}

	return
}
