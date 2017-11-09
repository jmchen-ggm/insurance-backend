package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"encoding/json"
)

func GetListArticle(bbReq protocol.BBReq) []byte {
	var listArticleRequest protocol.BBListArticleRequest
	json.Unmarshal(bbReq.Body, &listArticleRequest)
	articleList := database.GetListArticle(listArticleRequest.StartIndex, listArticleRequest.PageSize)
	log.Info("req %d %d %d", listArticleRequest.StartIndex, listArticleRequest.PageSize, articleList.len)
	var response protocol.BBListArticleResponse
	response.ArticleList = articleList
	responseBytes, _ := json.Marshal(response)
	return responseBytes
}
