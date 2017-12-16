package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/time"
	"com/bbinsurance/userserver/constants"
	"com/bbinsurance/userserver/database"
	"com/bbinsurance/userserver/protocol"
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
	user, err := database.GetUserByUsername(bbLoginRequest.Username)
	log.Info("FunLogin: %s userId: %d", bbReq.Body, user.Id)
	if err != nil {
		return nil, webcommon.ResponseCodeRequestInvalid, "Invalid Username: " + bbLoginRequest.Username
	} else {
		password, err := database.GetPasswordByUserId(user.Id)
		if err != nil {
			return nil, webcommon.ResponseCodeRequestInvalid, "Invalid Username: " + bbLoginRequest.Username
		} else {
			if password.PasswordMD5 == bbLoginRequest.PasswordMD5 {
				password.LastLoginToken = uuid.NewV4().String()
				password.Timestamp = time.GetTimestampInMilli()
				err := database.UpdateToken(password)
				if err != nil {
					return nil, webcommon.ResponseCodeRequestInvalid, "Invalid Password"
				} else {
					var bbLogicResponse protocol.BBLoginResponse
					bbLogicResponse.UserInfo = user
					bbLogicResponse.Token = password.LastLoginToken
					responseBytes, _ := json.Marshal(bbLogicResponse)
					return responseBytes, webcommon.ResponseCodeSuccess, ""
				}
			} else {
				return nil, webcommon.ResponseCodeRequestInvalid, "Invalid Password"
			}
		}
	}
}

func FunGetUser(bbReq webcommon.BBReq) ([]byte, int, string) {
	var bbGetUserRequest protocol.BBGetUserRequest
	json.Unmarshal(bbReq.Body, &bbGetUserRequest)
	user, err := database.GetUser(bbGetUserRequest.UserId)
	if err != nil {
		return nil, webcommon.ResponseCodeServerError, "Server Error"
	} else {
		var bbGetUserResponse protocol.BBGetUserResponse
		bbGetUserResponse.User = user
		responseBytes, _ := json.Marshal(bbGetUserResponse)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	}
}

func FunBatchGetUser(bbReq webcommon.BBReq) ([]byte, int, string) {
	var bbLoginRequest protocol.BBLoginRequest
	json.Unmarshal(bbReq.Body, &bbLoginRequest)
	return nil, 0, ""
}

func FunGetAllUser(bbReq webcommon.BBReq) ([]byte, int, string) {
	var batchGetUserResponse protocol.BBBatchGetUserResponse
	userList, err := database.GetAllUserList()
	if err != nil {
		return nil, webcommon.ResponseCodeServerError, "Server Error"
	} else {
		batchGetUserResponse.UserList = userList
		responseBytes, _ := json.Marshal(batchGetUserResponse)
		return responseBytes, webcommon.ResponseCodeSuccess, ""
	}

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
		user, err := database.GetUserByUsername(username)
		if err == nil {
			log.Error("Invalid User duplicated username %s", username)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeRequestInvalid, "Duplicated username: "+username)
			return
		}
		user.Username = username
		user.Nickname = request.FormValue("nickname")
		user.PhoneNumber = request.FormValue("phoneNumber")
		password := request.FormValue("password")
		usernameMD5 := util.MD5(user.Username)
		passwordMD5 := util.MD5(password)
		user.ThumbUrl = fmt.Sprintf("img/users/%s.png", usernameMD5)
		log.Info("CreateUser: %s file: %s", util.ObjToString(user), fileHandler.Header)
		savePath := constants.STATIC_FOLDER + "/" + user.ThumbUrl
		fis, err := util.FileCreate(savePath)
		defer fis.Close()
		if err != nil {
			log.Error("Save File Err %s", err)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeServerError, "Save File Error")
		} else {
			io.Copy(fis, file)
			id, _ := database.InsertUser(user)
			var passwordObj protocol.Password
			passwordObj.UserId = id
			passwordObj.PasswordMD5 = passwordMD5
			passwordObj.LastLoginToken = ""
			passwordObj.Timestamp = time.GetTimestampInMilli()
			log.Info("Create Password: %s", util.ObjToString(passwordObj))
			database.InsertPassword(passwordObj)
			webcommon.HandleSuccessResponse(writer, bbReq, nil)
		}
	}
}
