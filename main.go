package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type cache map[string][]byte

var (
	dbpool    sql.DB
	userCache = cache{}
)

func init() {
	db, err := sql.Open("sqlite3", "./plank.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userCache.grabUsers(db)
}

func main() {
	http.HandleFunc("/auth/logn", login)
	http.ListenAndServe(":9000", nil)
}
