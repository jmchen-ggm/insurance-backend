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
	log.Info("static folder %s", constants.STATIC_FOLDER)
	http.Handle("/", http.FileServer(http.Dir(constants.STATIC_FOLDER)))
	log.Info("listen %s port", constants.PORT)
	http.ListenAndServe(":"+constants.PORT, nil)
}
