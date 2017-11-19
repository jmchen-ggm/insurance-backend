package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"encoding/json"
)

func FunGetListComment(bbReq protocol.BBReq) ([]byte, int, string) {
	var listCommentRequest protocol.BBListCommentRequest
	json.Unmarshal(bbReq.Body, &listCommentRequest)
	commentList := database.GetListComment(listCommentRequest.StartIndex, listCommentRequest.PageSize)
	log.Info("req %d %d %d", listCommentRequest.StartIndex, listCommentRequest.PageSize, len(commentList))
	var response protocol.BBListCommentResponse
	response.CommentList = commentList
	responseBytes, _ := json.Marshal(response)
	return responseBytes, protocol.ResponseCodeSuccess, ""
}

func FunCreateComment(bbReq protocol.BBReq) ([]byte, int, string) {
	var createCommentRequest protocol.BBCreateCommentRequest
	json.Unmarshal(bbReq.Body, &createCommentRequest)
	id, err := database.InsertComment(createCommentRequest.Comment.Uin, createCommentRequest.Comment.Content,
		createCommentRequest.Comment.Score)
	log.Info("FuncCreateComment: %d", id)
	if err != nil {
		log.Error("FuncCreateComment %s", err)
		return nil, protocol.ResponseCodeServerError, "Create Comment Error"
	} else {
		var response protocol.BBCreateCommentResponse
		response.ServerId = id
		responseBytes, _ := json.Marshal(response)
		return responseBytes, protocol.ResponseCodeSuccess, ""
	}
}

func FunViewComment(bbReq protocol.BBReq) ([]byte, int, string) {
	var viewCommentRequest protocol.BBViewCommentRequest
	json.Unmarshal(bbReq.Body, &viewCommentRequest)
	database.UpdateCommentViewCount(viewCommentRequest.ServerId)
	log.Info("FuncCreateComment: %d", viewCommentRequest.ServerId)
	return nil, protocol.ResponseCodeSuccess, ""
}
