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


func GetListCompany(bbReq protocol.BBReq) []byte {
	var listCompanyRequest protocol.BBListCompanyRequest
	json.Unmarshal(bbReq.Body, &listCompanyRequest)
	companyList := database.GetListCompany(listCompanyRequest.StartIndex, listCompanyRequest.PageSize)
	log.Info("req %d %d %d", listCompanyRequest.StartIndex, listCompanyRequest.PageSize, len(companyList))
	var response protocol.BBListCompanyResponse
	response.CompanyList = companyList
	responseBytes, _ := json.Marshal(response)
	return responseBytes
}

func HandleCreateCompany(writer http.ResponseWriter, request *http.Request) {
	var bbReq protocol.BBReq
	bbReq.Bin.FunId = protocol.FuncCreateCompany
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
		name := request.FormValue("name")
		desc := request.FormValue("desc")

		log.Info("CreateCompany: name=%s desc=%s file=%s", name, desc, fileHandler.Header)

		id, err := database.InsertCompany(name, desc, "")
		if err != nil {
			log.Error("Invalid File %s", err)
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeServerError, "Insert Company Error")
			return
		}
		thumbUrl := fmt.Sprintf("img/companys/%d.png", id)
		database.UpdateCompanyThumbUrl(id, thumbUrl)
		savePath := constants.STATIC_FOLDER + "/" + thumbUrl
		fis, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
		defer fis.Close()
		if err != nil {
			log.Error("Save File Err %s", err)
			database.DeleteCompanyById(id)
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeServerError, "Save File Error")
		} else {
			log.Info("Save File success %s", savePath)
			io.Copy(fis, file)
			var response protocol.BBCreateCompanyResponse
			response.Id = id
			response.ThumbUrl = thumbUrl
			responseBytes, _ := json.Marshal(response)
			HandleSuccessResponse(writer, bbReq, responseBytes)
		}
	}
}