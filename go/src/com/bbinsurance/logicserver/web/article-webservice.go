package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/constants"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/time"
	"com/bbinsurance/util"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

func GetListArticle(bbReq protocol.BBReq) []byte {
	var listArticleRequest protocol.BBListArticleRequest
	json.Unmarshal(bbReq.Body, &listArticleRequest)
	articleList := database.GetListArticle(listArticleRequest.StartIndex, listArticleRequest.PageSize)
	log.Info("req %d %d %d", listArticleRequest.StartIndex, listArticleRequest.PageSize, len(articleList))
	var response protocol.BBListArticleResponse
	response.ArticleList = articleList
	responseBytes, _ := json.Marshal(response)
	return responseBytes
}

func HandleCreateArticle(writer http.ResponseWriter, request *http.Request) {
	var bbReq protocol.BBReq
	bbReq.Bin.FunId = protocol.FuncCreateArticle
	bbReq.Bin.URI = protocol.UriCreateData
	bbReq.Bin.SessionId = uuid.NewV4().String()
	bbReq.Bin.Timestamp = time.GetTimestamp()
	if request.Method != "POST" {
		log.Error("Invalid Request Method: %s Url: %s", request.Method, request.URL)
		HandleErrorResponse(writer, bbReq, protocol.ResponseCodeRequestInvalid, "Invalid Requst, Please Use Http POST")
		return
	} else {
		request.ParseMultipartForm(32 << 20)
		file, fileHandler, err := request.FormFile("uploadfile")
		defer file.Close()
		if err != nil {
			log.Error("Invalid File %s", err)
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeRequestInvalid, "Invalid Requst File")
			return
		}
		title := request.FormValue("title")
		desc := request.FormValue("desc")
		date := request.FormValue("date")
		url := request.FormValue("url")

		log.Info("CreateArticle: title=%s desc=%s date=%s url=%s file=%s", title, desc, date, url, fileHandler.Header)

		id, err := database.InsertArticle(title, desc, date, url, "")
		if err != nil {
			log.Error("Invalid File %s", err)
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeServerError, "Insert Article Error")
			return
		}
		thumbUrl := fmt.Sprintf("img/articles/%d.png", id)
		database.UpdateArticleThumbUrl(id, thumbUrl)
		savePath := constants.STATIC_FOLDER + "/" + thumbUrl
		fis, err := util.FileCreate(savePath)
		defer fis.Close()
		if err != nil {
			log.Error("Save File Err %s", err)
			database.DeleteArticleById(id)
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeServerError, "Save File Error")
		} else {
			log.Info("Save File success %s", savePath)
			io.Copy(fis, file)
			var response protocol.BBCreateArticleResponse
			response.Id = id
			response.ThumbUrl = thumbUrl
			responseBytes, _ := json.Marshal(response)
			HandleSuccessResponse(writer, bbReq, responseBytes)
		}
	}
}
