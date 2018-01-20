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
	"strconv"
)

func FunGetListCompany(bbReq webcommon.BBReq) ([]byte, int, string) {
	var listCompanyRequest protocol.BBListCompanyRequest
	json.Unmarshal(bbReq.Body, &listCompanyRequest)
	companyList := database.GetListCompany(listCompanyRequest.StartIndex, listCompanyRequest.PageSize)
	log.Info("req %d %d %d", listCompanyRequest.StartIndex, listCompanyRequest.PageSize, len(companyList))
	var response protocol.BBListCompanyResponse
	response.CompanyList = companyList
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}

func FunGetCompanyById(bbReq webcommon.BBReq) ([]byte, int, string) {
	var getCompanyRequest protocol.BBGetCompanyRequest
	json.Unmarshal(bbReq.Body, &getCompanyRequest)
	company := service.GetCompanyById(getCompanyRequest.Id)
	if company.Id == -1 {
		log.Error("FunGetCompanyById Err")
		return nil, webcommon.ResponseCodeServerError, "FunGetCompanyById Err"
	} else {
		var response protocol.BBGetCompanyResponse
		response.Company = company
		responseBytes, _ := json.Marshal(response)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	}
}

func FunCreateCompany(writer http.ResponseWriter, request *http.Request) {
	var bbReq webcommon.BBReq
	bbReq.Bin.FunId = webcommon.FuncCreateCompany
	bbReq.Bin.URI = webcommon.UriCreateData
	bbReq.Bin.SessionId = uuid.NewV4().String()
	bbReq.Bin.Timestamp = time.GetTimestamp()
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
		var company protocol.Company
		company.Name = request.FormValue("name")
		company.Desc = request.FormValue("desc")
		company.Flags, _ = strconv.ParseInt(request.FormValue("flags"), 16, 64)
		company.DetailData = request.FormValue("detailData")
		company.ThumbUrl = fmt.Sprintf("img/companys/%s.png", uuid.NewV4().String())

		savePath := constants.STATIC_FOLDER + "/" + company.ThumbUrl
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
		log.Info("CreateCompany: %s", util.ObjToString(company))
		company, err = database.InsertCompany(company)
		if err != nil {
			util.DeleteFile(savePath)
			log.Error("Insert data to db error %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Create Company Error")
		} else {
			var response protocol.BBCreateCompanyResponse
			response.Company = company
			responseBytes, _ := json.Marshal(response)
			webcommon.HandleSuccessResponse(writer, bbReq, responseBytes)
		}
	}
}
