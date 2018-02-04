package protocol

type User struct {
	Id        int64
	Username  string
	Nickname  string
	Timestamp int64
	ThumbUrl  string
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

type BBListUserRequest struct {
	StartIndex int
	PageSize   int
}

type BBListUserResponse struct {
	UserList []User
}

type BBLoginRequest struct {
	Username    string
	PasswordMD5 string
}

type BBLoginResponse struct {
	LoginCode int
	User      User
	Token     string
	Timestamp int64
}
