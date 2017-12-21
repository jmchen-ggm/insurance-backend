package main

import (
	"com/bbinsurance/homeserver/constants"
	"com/bbinsurance/log"
	"net/http"
)

func main() {
	// 初始化Logger
	log.InitLogging("bb-insurance.log")
	log.Info("Home Server Start")
	// 初始化变量
	constants.InitConstants()
	//初始Http路径监控
	http.Handle("/", http.FileServer(http.Dir(constants.STATIC_FOLDER)))
	http.ListenAndServe(":"+constants.PORT, nil)
}
