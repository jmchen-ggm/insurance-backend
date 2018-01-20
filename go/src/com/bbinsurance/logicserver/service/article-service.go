package service

import (
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/webcommon"
)

func GetListArticle(startIndex int, length int) []protocol.Article {
	articleList := database.GetListArticle(startIndex, length)
	for i := 0; i < len(articleList); i++ {
		articleList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(articleList[i].ThumbUrl)
	}
	return articleList
}
