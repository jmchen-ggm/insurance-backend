package web

import (
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"encoding/json"
)

func GetListArticle(bbReq protocol.BBReq) json.RawMessage {
	var listArticleRequest protocol.BBListArticleRequest
	json.Unmarshal(bbReq.Body, &listArticleRequest)
	articleList := database.GetListArticle(listArticleRequest.StartIndex, listArticleRequest.PageSize)
	var response protocol.BBListArticleResponse
	response.ArticleList = articleList
	var responseRawMessage json.RawMessage
	responseBytes, _ := json.Marshal(response)
	json.Unmarshal(responseBytes, &responseRawMessage)
	return responseRawMessage
}
