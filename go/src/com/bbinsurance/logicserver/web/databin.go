package web

import (
	"com/bbinsurance/log"
	"net/http"
	// "com/bbinsurance/logicserver/database"
)

func HandleDataBin(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		log.Error("Invalid Request Method: %s Url: %s", request.Method, request.URL)
	}
}
