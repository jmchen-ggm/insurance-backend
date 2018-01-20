package webcommon

type HttpFun struct {
	Id           int64
	FunId        int64
	Timestamp    int64
	ResponseSize int64
	UseTime      int64
	Uin          int64
}

type KvHttpFunRequest struct {
	HttpFun HttpFun
}

type KvHttpFunResponse struct {
}
