package service

import (
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
)

func GetListArticle(startIndex int, length int) []protocol.Article {
	articleList := database.GetListArticle(startIndex, length)
	return articleList
}
