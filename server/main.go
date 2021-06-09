package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/http-sqlite-rest/server/database/function"
	"github.com/sandjuarezg/http-sqlite-rest/server/database/user"
)

var db *sql.DB
var err error

func main() {
	function.SqlMigration()

	db, err = sql.Open("sqlite3", "./database/user.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/add", postAdd)
	http.HandleFunc("/show", getShow)

	fmt.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var addU function.User
	err = json.NewDecoder(r.Body).Decode(&addU)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = user.AddUser(db, addU)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode("Insert data successfully")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	users, err := user.ShowUser(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
