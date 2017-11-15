package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"encoding/json"
)

func FunGetListComment(bbReq protocol.BBReq) []byte {
	var listCommentRequest protocol.BBListCommentRequest
	json.Unmarshal(bbReq.Body, &listCommentRequest)
	commentList := database.GetListComment(listCommentRequest.StartIndex, listCommentRequest.PageSize)
	log.Info("req %d %d %d", listCommentRequest.StartIndex, listCommentRequest.PageSize, len(commentList))
	var response protocol.BBListCommentResponse
	response.CommentList = commentList
	responseBytes, _ := json.Marshal(response)
	return responseBytes
}

func FunCreateComment(bbReq protocol.BBReq) []byte {
	var createCommentRequest protocol.BBCreateCommentRequest
	json.Unmarshal(bbReq.Body, &createCommentRequest)
	id, err := database.InsertComment(createCommentRequest.Comment.Uin, createCommentRequest.Comment.Content,
		createCommentRequest.Comment.Score)
	log.Info("FuncCreateComment: %d", id)
	if err != nil {
		log.Error("FuncCreateComment %s", err)
	}
	var response protocol.BBCreateCommentResponse
	response.ServId = id
	responseBytes, _ := json.Marshal(response)
	return responseBytes
}

func FunViewComment(bbReq protocol.BBReq) []byte {
	var viewCommentRequest protocol.BBViewCommentRequest
	json.Unmarshal(bbReq.Body, &viewCommentRequest)
	database.UpdateCommentViewCount(viewCommentRequest.ServId)
	log.Info("FuncCreateComment: %d", viewCommentRequest.ServId)
	return nil
}
