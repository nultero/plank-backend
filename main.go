package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/auth/logn", login)
	http.ListenAndServe(":9000", nil)
}
