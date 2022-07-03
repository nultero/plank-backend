package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

const (
	msgLastRq     = "couldn't make sense of last request"
	msgMangldRq   = "mangled json on request"
	msgUserExists = "user already exists"
)

var (
	errLRq        = errors.New(msgLastRq)
	errMangldRq   = errors.New(msgMangldRq)
	errUserExists = errors.New(msgUserExists)
)

func parseUreqs(w http.ResponseWriter, rq *http.Request) (map[string]string, error) {
	bytes := make([]byte, 50)
	n, err := rq.Body.Read(bytes)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			fmt.Fprintln(w, msgLastRq)
			return nil, errLRq
		}
	}

	bytes = bytes[:n]
	mp := map[string]string{}
	err = json.Unmarshal(bytes, &mp)
	if err != nil {
		fmt.Fprintln(w, msgMangldRq)
		return nil, errMangldRq
	}

	return mp, nil
}

func createUser(w http.ResponseWriter, rq *http.Request) {
	mp, err := parseUreqs(w, rq)
	if err != nil {
		return
	}

	for k, v := range mp {
		if userCache.hasUser(k) {
			fmt.Fprintln(w, msgUserExists)
			return
		}

		fmt.Println(v)

	}

}

func login(w http.ResponseWriter, rq *http.Request) {

	mp, err := parseUreqs(w, rq)
	if err != nil {
		return
	}

	for k, v := range mp {
		if pw, ok := userCache[k]; ok {
			err = bcrypt.CompareHashAndPassword(pw, []byte(v))
			if err == nil {
				fmt.Fprintln(w, "match!")
				return
			}
		}

		fmt.Fprintln(w, "not a match")
		break
	}
}
