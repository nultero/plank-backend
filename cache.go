package main

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func (c cache) hasUser(userName string) bool {
	_, ok := c[userName]
	return ok
}

func (c cache) grabUsers(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		uname := ""
		pwd := ""
		err := rows.Scan(&uname, &pwd)
		if err != nil {
			return err
		}

		hp, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost-6)
		if err != nil {
			panic(err)
		}

		c[uname] = hp
	}
	return nil
}

func (c cache) populateUsers(db sql.DB) {

	// db.Exec()
}
