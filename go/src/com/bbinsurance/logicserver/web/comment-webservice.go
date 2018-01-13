package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/logicserver/service"
	"com/bbinsurance/time"
	"com/bbinsurance/webcommon"
	"encoding/json"
)

func FunGetListComment(bbReq webcommon.BBReq) ([]byte, int, string) {
	var listCommentRequest protocol.BBListCommentRequest
	json.Unmarshal(bbReq.Body, &listCommentRequest)
	commentList := service.GetListComment(bbReq.Header.Uin, listCommentRequest.StartIndex, listCommentRequest.PageSize)
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
	var err error
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
	var response protocol.BBViewCommentResponse
	response.Comment = service.ViewComment(viewCommentRequest.Id)
	if response.Comment.Id == -1 {
		return nil, webcommon.ResponseCodeServerError, "Not Found Comment"
	} else {
		responseBytes, _ := json.Marshal(response)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	}
}

func FunUpComment(bbReq webcommon.BBReq) ([]byte, int, string) {
	var upCommentRequest protocol.BBUpCommentRequest
	json.Unmarshal(bbReq.Body, &upCommentRequest)
	var response protocol.BBUpCommentResponse
	response.Comment = service.UpComment(upCommentRequest.CommentUp, upCommentRequest.IsUp)
	if response.Comment.Id == -1 {
		return nil, webcommon.ResponseCodeServerError, "Not Found Comment"
	} else {
		responseBytes, _ := json.Marshal(response)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	}
}

func FunReplyComment(bbReq webcommon.BBReq) ([]byte, int, string) {
	var replyCommentRequest protocol.BBReplyCommentRequest
	json.Unmarshal(bbReq.Body, &replyCommentRequest)
	var response protocol.BBReplyCommentResponse
	response.Comment = service.ReplyComment(replyCommentRequest.CommentReply)
	if response.Comment.Id == -1 {
		return nil, webcommon.ResponseCodeServerError, "Not Found Comment"
	} else {
		responseBytes, _ := json.Marshal(response)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	}
}

func FunGetListCommentReply(bbReq webcommon.BBReq) ([]byte, int, string) {
	var getCommentReplyListRequest protocol.BBGetCommentReplyListRequest
	json.Unmarshal(bbReq.Body, &getCommentReplyListRequest)
	var response protocol.BBGetCommentReplyListResponse
	response.CommentReplyList = service.GetListCommentReply(getCommentReplyListRequest.CommentId)
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}
