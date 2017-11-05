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

	//初始Http路径监控
	http.Handle("/static/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/create/article/", web.HandleCreateArticle)
	http.ListenAndServe(":8081", nil)

	// database.InsertArticle("真心不错", "不错不错", "http://", "http://")
	// var articleList *list.List = database.GetAllArticle()
	// var articleSlince = make([]database.Article, articleList.Len())

	// for e := articleList.Front(); e != nil; e = e.Next() {
	// 	article := e.Value.(database.Article)
	// 	articleSlince = append(articleSlince, article)
	// }

	// listData, _ := json.Marshal(articleSlince)
	// fmt.Println(string(listData))

	// for e := articleList.Front(); e != nil; e = e.Next() {
	// 	article := e.Value.(database.Article)
	// 	data, _ := json.Marshal(article)
	// 	fmt.Println(string(data))
	// }
}
