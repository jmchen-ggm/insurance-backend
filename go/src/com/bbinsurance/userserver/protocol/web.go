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
	PasswordMd5    string
	LastLoginToken string
	Timestamp      int64
}
