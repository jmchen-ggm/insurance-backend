package service

import (
	"com/bbinsurance/time"
	"com/bbinsurance/userserver/database"
	"com/bbinsurance/userserver/protocol"
	"com/bbinsurance/webcommon"
	"github.com/satori/go.uuid"
)

func Login(request protocol.BBLoginRequest) protocol.BBLoginResponse {
	var response protocol.BBLoginResponse
	user := GetUserByUsername(request.Username)
	if user.Id < 0 {
		response.LoginCode = protocol.LoginCodeNotFoundUser
		return response
	}
	password := database.GetPasswordByUserId(user.Id)
	if password.PasswordMD5 != request.PasswordMD5 {
		response.LoginCode = protocol.LoginCodePasswordErr
		return response
	}
	password.LastLoginToken = uuid.NewV4().String()
	password.Timestamp = time.GetTimestampInMilli()
	password = database.UpdateToken(password)
	if password.UserId < 0 {
		response.LoginCode = protocol.LoginCodeUpdateTokenErr
		return response
	}
	response.User = user
	response.Token = password.LastLoginToken
	response.Timestamp = password.Timestamp
	return response
}

func ListUser(request protocol.BBListUserRequest) protocol.BBListUserResponse {
	userList := database.ListUser(request.StartIndex, request.PageSize)
	for i := 0; i < len(userList); i++ {
		userList[i].ThumbUrl = webcommon.GenerateImgFileServerUrl(userList[i].ThumbUrl)
	}
	var response protocol.BBListUserResponse
	response.UserList = userList
	return response
}

func GetUserById(id int64) protocol.User {
	user := database.GetUserById(id)
	user.ThumbUrl = webcommon.GenerateImgFileServerUrl(user.ThumbUrl)
	return user
}

func GetUserByUsername(username string) protocol.User {
	user := database.GetUserByUsername(username)
	user.ThumbUrl = webcommon.GenerateImgFileServerUrl(user.ThumbUrl)
	return user
}
