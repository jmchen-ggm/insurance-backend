package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/time"
	"com/bbinsurance/userserver/constants"
	"com/bbinsurance/userserver/database"
	"com/bbinsurance/userserver/protocol"
	"com/bbinsurance/util"
	"com/bbinsurance/webcommon"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

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
		user, _ := database.GetUserByUsername(username)
		if user.Id >= 0 {
			log.Error("Invalid User duplicated username %s", username)
			webcommon.HandleErrorResponse(writer, bbReq, webcommon.ResponseCodeRequestInvalid, "Duplicated username")
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
