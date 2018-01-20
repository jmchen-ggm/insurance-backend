package webcommon

type HttpFun struct {
	Id           int64
	FunId        int
	Timestamp    int64
	ResponseSize int
	UseTime      int64
	Uin          int64
}

type KvHttpFunRequest struct {
	HttpFun HttpFun
}

type KvHttpFunResponse struct {
}
