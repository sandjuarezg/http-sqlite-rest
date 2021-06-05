package user

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/http-sqlite-rest/server/database/function"
)

func AddUser(db *sql.DB, user function.User) (err error) {
	smt, err := db.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
	if err != nil {
		return
	}
	defer smt.Close()

	_, err = smt.Exec(user.Name, user.Pass)
	if err != nil {
		return
	}

	return
}

func ShowUser(db *sql.DB) (users []function.User, err error) {
	rows, err := db.Query("SELECT id, name, password FROM users")
	if err != nil {
		return
	}
	defer rows.Close()

	var user function.User

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Pass)
		if err != nil {
			return
		}
		users = append(users, user)
	}

	return
}
