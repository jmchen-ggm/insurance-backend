package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/time"
	"com/bbinsurance/webcommon"
	"encoding/json"
)

func FunGetListComment(bbReq webcommon.BBReq) ([]byte, int, string) {
	var listCommentRequest protocol.BBListCommentRequest
	json.Unmarshal(bbReq.Body, &listCommentRequest)
	commentList := database.GetListComment(listCommentRequest.StartIndex, listCommentRequest.PageSize)
	log.Info("req %d %d %d", listCommentRequest.StartIndex, listCommentRequest.PageSize, len(commentList))
	var response protocol.BBListCommentResponse
	response.CommentList = commentList
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}

func FunCreateComment(bbReq webcommon.BBReq) ([]byte, int, string) {
	var createCommentRequest protocol.BBCreateCommentRequest
	json.Unmarshal(bbReq.Body, &createCommentRequest)
	createCommentRequest.Comment.Timestamp = time.GetTimestampInMilli()
	createCommentRequest.Comment.Flags = protocol.CREATED
	log.Info("Create Comment %s", time.GetCurrentTimeFormat(time.DATE_TIME_FMT))
	createCommentRequest.Comment, err = database.InsertComment(createCommentRequest.Comment)
	if err != nil {
		log.Error("FuncCreateComment %s", err)
		return nil, webcommon.ResponseCodeServerError, "Create Comment Error"
	} else {
		var response protocol.BBCreateCommentResponse
		response.Comment = createCommentRequest.Comment
		responseBytes, _ := json.Marshal(response)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	}
}

func FunViewComment(bbReq webcommon.BBReq) ([]byte, int, string) {
	var viewCommentRequest protocol.BBViewCommentRequest
	json.Unmarshal(bbReq.Body, &viewCommentRequest)
	database.UpdateCommentViewCount(viewCommentRequest.ServerId)
	log.Info("FuncCreateComment: %d", viewCommentRequest.ServerId)
	return nil, webcommon.ResponseCodeSuccess, ""
}
