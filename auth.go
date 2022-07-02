package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var userCache = map[string][]byte{}

// dumb prototype of a db pwd
func init() {
	mc, err := bcrypt.GenerateFromPassword([]byte("turbulence"), bcrypt.DefaultCost-6)
	if err != nil {
		panic(err)
	}

	userCache["jackson"] = mc
}

func login(w http.ResponseWriter, rq *http.Request) {
	bytes := make([]byte, 50)
	n, err := rq.Body.Read(bytes)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			fmt.Fprintln(w, "couldn't make sense of last request")
			return
		}
	}

	bytes = bytes[:n]
	mp := map[string]string{}
	err = json.Unmarshal(bytes, &mp)
	if err != nil {
		fmt.Fprintln(w, "mangled json on request")
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
