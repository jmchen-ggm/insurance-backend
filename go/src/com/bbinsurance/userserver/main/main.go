package main

import (
	"com/bbinsurance/log"
	"com/bbinsurance/userserver/constants"
	"com/bbinsurance/userserver/database"
	"com/bbinsurance/userserver/web"
	"net/http"
)

func main() {

	// 初始化Logger
	log.InitLogging("bb-insurance.log")
	log.Info("User Server Start")

	// 初始化变量
	constants.InitConstants()

	//初始化DB
	database.InitDB()

	http.HandleFunc("/create/user", web.FunCreateUser)

	//handle请求数据接口
	web.FunInitDataBin()
	http.HandleFunc("/data-bin", web.FunHandleDataBin)

	log.Info("listen %s port", constants.PORT)
	http.ListenAndServe(":"+constants.PORT, nil)
}
