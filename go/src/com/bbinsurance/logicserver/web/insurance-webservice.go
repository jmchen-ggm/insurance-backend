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

func FunGetListInsurance(bbReq webcommon.BBReq) ([]byte, int, string) {
	var listInsuranceRequest protocol.BBListInsuranceRequest
	json.Unmarshal(bbReq.Body, &listInsuranceRequest)
	insuranceList := service.GetListInsurance(listInsuranceRequest.StartIndex, listInsuranceRequest.PageSize)
	log.Info("req %d %d %d", listInsuranceRequest.StartIndex, listInsuranceRequest.PageSize, len(insuranceList))
	var response protocol.BBListInsuranceResponse
	response.InsuranceList = insuranceList
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}

func FunGetListInsuranceType(bbReq webcommon.BBReq) ([]byte, int, string) {
	var listInsuranceTypeRequest protocol.BBListInsuranceTypeRequest
	json.Unmarshal(bbReq.Body, &listInsuranceTypeRequest)
	insuranceTypeList := database.GetListInsuranceType(listInsuranceTypeRequest.StartIndex, listInsuranceTypeRequest.PageSize)
	log.Info("req %d %d %d", listInsuranceTypeRequest.StartIndex, listInsuranceTypeRequest.PageSize, len(insuranceTypeList))
	var response protocol.BBListInsuranceTypeResponse
	response.InsuranceTypeList = insuranceTypeList
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}

func FunGetInsuranceTypeById(bbReq webcommon.BBReq) ([]byte, int, string) {
	var getInsuranceTypeRequest protocol.BBGetInsuranceTypeRequest
	json.Unmarshal(bbReq.Body, &getInsuranceTypeRequest)
	insuranceType := service.GetInsuranceTypeById(getInsuranceTypeRequest.Id)
	if insuranceType.Id == -1 {
		log.Error("FunGetInsuranceTypeById Err")
		return nil, webcommon.ResponseCodeServerError, "FunGetInsuranceTypeById Err"
	} else {
		var response protocol.BBGetInsuranceTypeResponse
		response.InsuranceType = insuranceType
		responseBytes, _ := json.Marshal(response)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	}
}

func FunCreateInsuranceType(writer http.ResponseWriter, request *http.Request) {
	var bbReq webcommon.BBReq
	bbReq.Bin.FunId = webcommon.FuncCreateInsuranceType
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
		var insuranceType protocol.InsuranceType
		insuranceType.Name = request.FormValue("name")
		insuranceType.Desc = request.FormValue("desc")
		insuranceType.ThumbUrl = fmt.Sprintf("img/insuranceTypes/%s.png", uuid.NewV4().String())

		savePath := constants.STATIC_FOLDER + "/" + insuranceType.ThumbUrl
		log.Info("try to save file to path %s %s", savePath, fileHandler.Header)
		fis, err := util.FileCreate(savePath)
		defer fis.Close()
		if err != nil {
			log.Error("Create File Err %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Create File Error")
			return
		}
		_, err = io.Copy(fis, file)
		if err != nil {
			log.Error("Copy File Err %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Copy File Error")
			return
		}
		log.Info("CreateInsuranceType: %s", util.ObjToString(insuranceType))
		insuranceType, err = database.InsertInsuranceType(insuranceType)
		if err != nil {
			util.DeleteFile(savePath)
			log.Error("Insert data to db error %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Insert Insurance Error")
		} else {
			var response protocol.BBCreateInsuranceTypeResponse
			response.InsuranceType = insuranceType
			responseBytes, _ := json.Marshal(response)
			webcommon.HandleSuccessResponse(writer, bbReq, responseBytes)
		}
	}
}

func FunCreateInsurance(writer http.ResponseWriter, request *http.Request) {
	var bbReq webcommon.BBReq
	bbReq.Bin.FunId = webcommon.FuncCreateInsurance
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
		var insurance protocol.Insurance
		insurance.Name = request.FormValue("name")
		insurance.Desc = request.FormValue("desc")
		insurance.InsuranceTypeId, _ = strconv.ParseInt(request.FormValue("insuranceTypeId"), 10, 64)
		insurance.CompanyId, _ = strconv.ParseInt(request.FormValue("companyId"), 10, 64)
		insurance.AgeFrom, _ = strconv.Atoi(request.FormValue("ageFrom"))
		insurance.AgeTo, _ = strconv.Atoi(request.FormValue("ageTo"))
		insurance.AnnualCompensation, _ = strconv.Atoi(request.FormValue("annualCompensation"))
		insurance.AnnualPremium, _ = strconv.Atoi(request.FormValue("annualPremium"))
		insurance.Flags, _ = strconv.ParseInt(request.FormValue("flags"), 10, 64)
		insurance.DetailData = request.FormValue("detailData")
		insurance.ThumbUrl = fmt.Sprintf("img/insurances/%s.png", uuid.NewV4().String())

		insurance.InsuranceTypeName = database.GetInsuranceTypeNameById(insurance.InsuranceTypeId)
		if util.IsEmpty(insurance.InsuranceTypeName) {
			log.Error("Not Found Insurance Type Name %d", insurance.InsuranceTypeId)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeRequestInvalid, "Not Found Insurance Type Name")
			return
		}
		insurance.InsuranceTypeName = database.GetInsuranceTypeNameById(insurance.InsuranceTypeId)
		if util.IsEmpty(insurance.InsuranceTypeName) {
			log.Error("Not Found InsuranceType Name %d", insurance.InsuranceTypeId)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeRequestInvalid, "Not Found Insurance Type Name")
			return
		}

		savePath := constants.STATIC_FOLDER + "/" + insurance.ThumbUrl
		log.Info("try to save file to path %s %s", savePath, fileHandler.Header)
		fis, err := util.FileCreate(savePath)
		defer fis.Close()
		if err != nil {
			log.Error("Create File Err %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Create File Error")
			return
		}
		_, err = io.Copy(fis, file)
		if err != nil {
			log.Error("Copy File Err %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Copy File Error")
			return
		}
		log.Info("CreateInsurance: %s", util.ObjToString(insurance))
		insurance, err = database.InsertInsurance(insurance)
		if err != nil {
			util.DeleteFile(savePath)
			log.Error("Insert data to db error %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Insert Insurance Error")
		} else {
			var response protocol.BBCreateInsuranceResponse
			response.Insurance = insurance
			responseBytes, _ := json.Marshal(response)
			webcommon.HandleSuccessResponse(writer, bbReq, responseBytes)
		}
	}
}
