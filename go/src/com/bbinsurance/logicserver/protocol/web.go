package protocol

import (
	"encoding/json"
)

type BBBin struct {
	FunId     int
	URI       string
	SessionId string
	Timestamp int64
}

type BBReqHeader struct {
	Username string
	Token    string
}

type BBRespHeader struct {
	Username     string
	ResponseCode int
	ErrMsg       string
}

type BBResp struct {
	Bin    BBBin
	Header BBRespHeader
	Body   json.RawMessage
}

type BBReq struct {
	Bin    BBBin
	Header BBReqHeader
	Body   json.RawMessage
}

type BBCreateArticleResponse struct {
	Id       int64
	ThumbUrl string
}

type BBListArticleRequest struct {
	StartIndex int
	PageSize   int
}

type Article struct {
	Id       int
	Title    string
	Desc     string
	Url      string
	ThumbUrl string
}

type BBListArticleResponse struct {
	ArticleList []Article
}
