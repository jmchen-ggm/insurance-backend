package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/root/github/insurance-file/static/")))
	http.ListenAndServe(":8082", nil)
}
