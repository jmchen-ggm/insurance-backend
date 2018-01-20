package main

import (
	"com/bbinsurance/kvserver/constants"
	"com/bbinsurance/kvserver/database"
	"com/bbinsurance/kvserver/web"
	"com/bbinsurance/log"
	"net/http"
)

func main() {
	// 初始化Logger
	log.InitLogging("bb-insurance.log")
	log.Info("KV Server Start")

	// 初始化变量
	constants.InitConstants()

	//初始化DB
	database.InitDB()

	http.HandleFunc("/data-bin", web.FunHandleDataBin)
	log.Info("listen %s port", constants.PORT)
	http.ListenAndServe(":"+constants.PORT, nil)
}
