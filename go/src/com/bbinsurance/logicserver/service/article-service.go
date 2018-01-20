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

func ViewArticle(id int64) protocol.Article {
	article := database.GetArticleById(id)
	article.ThumbUrl = webcommon.GenerateImgFileServerUrl(article.ThumbUrl)
	if article.Id != -1 {
		article.ViewCount++
		database.UpdateArticleViewCount(id, article.ViewCount)
	}
	return article
}
