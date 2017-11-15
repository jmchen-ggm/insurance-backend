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

	//初始Http路径监控
	http.Handle("/", http.FileServer(http.Dir(constants.STATIC_FOLDER)))

	http.HandleFunc("/create/article", web.FunCreateArticle)
	http.HandleFunc("/create/company", web.FunCreateCompany)
	http.HandleFunc("/create/insurance", web.FunCreateInsurance)

	//handle请求数据接口
	web.FunInitDataBin()
	http.HandleFunc("/data-bin", web.FunHandleDataBin)

	log.Info("listen %s port", constants.PORT)
	http.ListenAndServe(":"+constants.PORT, nil)
}
