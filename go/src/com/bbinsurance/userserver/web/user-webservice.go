package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/time"
	"com/bbinsurance/userserver/constants"
	"com/bbinsurance/userserver/database"
	"com/bbinsurance/userserver/protocol"
	"com/bbinsurance/userserver/service"
	"com/bbinsurance/util"
	"com/bbinsurance/webcommon"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

func FunLogin(bbReq webcommon.BBReq) ([]byte, int, string) {
	var bbLoginRequest protocol.BBLoginRequest
	json.Unmarshal(bbReq.Body, &bbLoginRequest)
	bbLogicResponse := service.Login(bbLoginRequest)
	responseBytes, _ := json.Marshal(bbLogicResponse)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}

func FunGetUser(bbReq webcommon.BBReq) ([]byte, int, string) {
	var bbGetUserRequest protocol.BBGetUserRequest
	json.Unmarshal(bbReq.Body, &bbGetUserRequest)
	user := service.GetUserById(bbGetUserRequest.UserId)
	if user.Id < 0 {
		return nil, webcommon.ResponseCodeRequestInvalid, "Not Found User"
	} else {
		var bbGetUserResponse protocol.BBGetUserResponse
		bbGetUserResponse.User = user
		responseBytes, _ := json.Marshal(bbGetUserResponse)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	}
}

func FunListUser(bbReq webcommon.BBReq) ([]byte, int, string) {
	var listUserRequest protocol.BBListUserRequest
	json.Unmarshal(bbReq.Body, &listUserRequest)
	response := service.ListUser(listUserRequest)
	responseBytes, _ := json.Marshal(response)
	return responseBytes, webcommon.ResponseCodeSuccess, ""
}

func FunCreateUser(writer http.ResponseWriter, request *http.Request) {
	var bbReq webcommon.BBReq
	bbReq.Bin.FunId = webcommon.FuncRegisterUser
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
		username := request.FormValue("username")
		user := service.GetUserByUsername(username)
		if user.Id >= 0 {
			log.Error("Invalid User duplicated username %s", username)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeRequestInvalid, "Duplicated username: "+username)
			return
		}
		user.Username = username
		user.Nickname = request.FormValue("nickname")
		password := request.FormValue("password")
		usernameMD5 := util.MD5(user.Username)
		passwordMD5 := util.MD5(password)
		user.ThumbUrl = fmt.Sprintf("img/users/%s.png", usernameMD5)
		user.Timestamp = time.GetTimestampInMilli()
		log.Info("CreateUser: %s file: %s", util.ObjToString(user), fileHandler.Header)
		savePath := constants.STATIC_FOLDER + "/" + user.ThumbUrl
		fis, err := util.FileCreate(savePath)
		defer fis.Close()
		if err != nil {
			log.Error("Save File Err %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Save File Error")
		} else {
			io.Copy(fis, file)
			user = database.InsertUser(user)
			if user.Id >= 0 {
				var passwordObj protocol.Password
				passwordObj.UserId = user.Id
				passwordObj.PasswordMD5 = passwordMD5
				passwordObj.LastLoginToken = ""
				passwordObj.Timestamp = time.GetTimestampInMilli()
				log.Info("Create Password: %s", util.ObjToString(passwordObj))
				database.InsertPassword(passwordObj)
				webcommon.HandleSuccessResponse(writer, bbReq, nil)
			} else {
				webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeRequestInvalid, "Insert User Error")
			}
		}
	}
}
