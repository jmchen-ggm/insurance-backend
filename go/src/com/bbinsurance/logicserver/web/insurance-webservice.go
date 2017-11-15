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
	"strconv"
)

func FunGetListInsurance(bbReq protocol.BBReq) []byte {
	var listInsuranceRequest protocol.BBListInsuranceRequest
	json.Unmarshal(bbReq.Body, &listInsuranceRequest)
	insuranceList := database.GetListInsurance(listInsuranceRequest.StartIndex, listInsuranceRequest.PageSize)
	log.Info("req %d %d %d", listInsuranceRequest.StartIndex, listInsuranceRequest.PageSize, len(insuranceList))
	var response protocol.BBListInsuranceResponse
	response.InsuranceList = insuranceList
	responseBytes, _ := json.Marshal(response)
	return responseBytes
}

func FunCreateInsurance(writer http.ResponseWriter, request *http.Request) {
	var bbReq protocol.BBReq
	bbReq.Bin.FunId = protocol.FuncCreateInsurance
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
		nameZHCN := request.FormValue("nameZHCN")
		nameEN := request.FormValue("nameEN")
		desc := request.FormValue("desc")
		companyIdStr := request.FormValue("companyId")
		companyId, _ := strconv.Atoi(companyIdStr)

		log.Info("CreateInsurance: nameZHCN=%s nameEN=%s desc=%s companyId=%d file=%s", nameZHCN, nameEN, desc, companyId, fileHandler.Header)

		id, err := database.InsertInsurance(nameZHCN, nameEN, desc, 0, companyId, "")
		if err != nil {
			log.Error("Invalid File %s", err)
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeServerError, "Insert Insurance Error")
			return
		}
		thumbUrl := fmt.Sprintf("img/insurances/%d.png", id)
		database.UpdateInsuranceThumbUrl(id, thumbUrl)
		savePath := constants.STATIC_FOLDER + "/" + thumbUrl
		fis, err := util.FileCreate(savePath)
		defer fis.Close()
		if err != nil {
			log.Error("Save File Err %s", err)
			database.DeleteInsuranceById(id)
			HandleErrorResponse(writer, bbReq, protocol.ResponseCodeServerError, "Save File Error")
		} else {
			log.Info("Save File success %s", savePath)
			io.Copy(fis, file)
			var response protocol.BBCreateInsuranceResponse
			response.Id = id
			response.ThumbUrl = thumbUrl
			responseBytes, _ := json.Marshal(response)
			HandleSuccessResponse(writer, bbReq, responseBytes)
		}
	}
}
