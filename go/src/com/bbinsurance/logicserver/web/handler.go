package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/constants"
	"com/bbinsurance/logicserver/database"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/time"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
	"os"
)

func HandleCreateArticle(writer http.ResponseWriter, request *http.Request) {
	var bbReq protocol.BBReq
	bbReq.Bin.FunId = protocol.FuncCreateArticle
	bbReq.Bin.URI = protocol.UriCreateArticle
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
		url := request.FormValue("url")

		log.Info("CreateArticle: title=%s desc=%s url=%s file=%s", title, desc, url, fileHandler.Header)

		id, err := database.InsertArticle(title, desc, url, "")
		if err != nil {
			log.Error("Invalid File %s", err)
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeServerError, "Insert Article Error")
			return
		}
		thumbUrl := fmt.Sprint(id) + ".png"
		database.UpdateArticleThumbUrl(id, thumbUrl)
		savePath := constants.STATIC_FOLDER + "/img/articles/" + thumbUrl
		fis, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Error("Save File Err %s", err)
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeServerError, "Save File Error")
			return
		} else {
			log.Info("Save File success %s", savePath)
		}
		defer fis.Close()
		io.Copy(fis, file)
		var response protocol.BBCreateArticleResponse
		response.Id = id
		response.ThumbUrl = thumbUrl
		responseBytes, _ := json.Marshal(response)
		HandleSuccessResponse(writer, bbReq, responseBytes)
	}
}
