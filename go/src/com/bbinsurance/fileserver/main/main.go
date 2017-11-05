package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir("/Users/jiaminchen/")))
	http.ListenAndServe(":8082", nil)
}
