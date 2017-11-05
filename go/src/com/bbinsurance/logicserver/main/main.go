package main

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/web"
	"net/http"
	// "container/list"
	// "encoding/json"
	// "fmt"
)

func main() {
	// 初始化Logger
	log.InitLogging("bb-insurance.log")
	log.Info("Server Start")

	//初始化DB
	database.InitDB()

	log.Info("handle /create/article/")
	//初始Http路径监控
	http.HandleFunc("/create/article/", web.HandleCreateArticle)

	log.Info("listen 8081 port")
	http.ListenAndServe(":8081", nil)
}
