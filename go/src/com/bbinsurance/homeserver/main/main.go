package main

import (
	"com/bbinsurance/homeserver/constants"
	"net/http"
)

func main() {
	//初始Http路径监控
	http.Handle("/", http.FileServer(http.Dir(constants.STATIC_FOLDER)))
	http.ListenAndServe(":"+constants.PORT, nil)
}
