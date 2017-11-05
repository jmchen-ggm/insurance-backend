package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("~/github/insurance-file/")))
	http.ListenAndServe(":8082", nil)
}
