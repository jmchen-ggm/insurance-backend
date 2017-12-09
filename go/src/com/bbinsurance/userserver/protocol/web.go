package protocol

type User struct {
	Id          int64
	Username    string
	Nickname    string
	PhoneNumber string
	Timestamp   int64
	ThumbUrl    string
}

type Password struct {
	UserId         int64
	PasswordMD5    string
	LastLoginToken string
	Timestamp      int64
}

type BBGetUserRequest struct {
	UserId int64
}

type BBGetUserResponse struct {
	User User
}

type BBBatchGetUserRequest struct {
	UserIdList []int64
}

type BBBatchGetUserResponse struct {
	UserList []User
}

type BBLoginRequest struct {
	Username    string
	PasswordMD5 string
}

type BBLoginResponse struct {
	UserInfo User
	Token    string
}
