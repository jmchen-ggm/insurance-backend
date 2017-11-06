package main

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/constants"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/web"
	"net/http"
)

func main() {
	// 初始化Logger
	log.InitLogging("bb-insurance.log")
	log.Info("Server Start")

	// 初始化变量
	constants.InitConstants()

	//初始化DB
	database.InitDB()

	log.Info("handle /create/article/")
	//初始Http路径监控
	http.HandleFunc("/create/article/", web.HandleCreateArticle)

	//handle请求数据接口
	http.HandleFunc("/data-bin/", web.HandleDataBin)

	log.Info("listen %s port", constants.PORT)
	http.ListenAndServe(":"+constants.PORT, nil)
}
