package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/constants"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/logicserver/service"
	"com/bbinsurance/time"
	"com/bbinsurance/util"
	"com/bbinsurance/webcommon"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

func FunGetListArticle(bbReq webcommon.BBReq) ([]byte, int, string) {
	var listArticleRequest protocol.BBListArticleRequest
	json.Unmarshal(bbReq.Body, &listArticleRequest)
	articleList := service.GetListArticle(listArticleRequest.StartIndex, listArticleRequest.PageSize)
	var response protocol.BBListArticleResponse
	response.ArticleList = articleList
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}

func FunCreateArticle(writer http.ResponseWriter, request *http.Request) {
	var bbReq webcommon.BBReq
	bbReq.Bin.FunId = webcommon.FuncCreateArticle
	bbReq.Bin.URI = webcommon.UriCreateData
	bbReq.Bin.SessionId = uuid.NewV4().String()
	bbReq.Bin.Timestamp = time.GetTimestampInMilli()
	if request.Method != "POST" {
		log.Error("Invalid Request Method: %s Url: %s", request.Method, request.URL)
		webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeRequestInvalid, "Invalid Requst, Please Use Http POST")
		return
	} else {
		request.ParseMultipartForm(32 << 20)
		file, fileHandler, err := request.FormFile("uploadfile")
		defer file.Close()
		if err != nil {
			log.Error("Invalid File %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeRequestInvalid, "Invalid Requst File")
			return
		}
		var article protocol.Article

		article.Title = request.FormValue("title")
		article.Desc = request.FormValue("desc")
		article.Date = request.FormValue("date")
		article.Url = request.FormValue("url")
		article.Timestamp = time.GetTimestampInMilli()
		aritcle.ThumbUrl = fmt.Sprintf("img/articles/%s.jpg", article.Title)

		savePath := constants.STATIC_FOLDER + "/" + aritcle.ThumbUrl
		log.Info("try to save file to path %s %s", savePath, fileHandler.Header)
		fis, err := util.FileCreate(savePath)
		defer fis.Close()
		if err != nil {
			log.Error("Save File Err %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Save File Error")
			return
		}

		_, err = io.Copy(fis, file)
		if err != nil {
			log.Error("Copy File Err %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Copy File Error")
			return
		}
		log.Info("CreateArticle: %s", util.ObjToString(article))
		aritcle, err := database.InsertArticle(aritcle)
		if err != nil {
			util.DeleteFile(savePath)
			log.Error("Insert data to db error %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Create Article Error")
		} else {
			var response protocol.BBCreateArticleResponse
			response.Article = aritcle
			responseBytes, _ := json.Marshal(response)
			webcommon.HandleSuccessResponse(writer, bbReq, responseBytes)
		}
	}
}
