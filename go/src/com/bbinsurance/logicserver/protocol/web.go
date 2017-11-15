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
	Uin      int64
	Username string
	Token    string
}

type BBRespHeader struct {
	Uin          int64
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
	Id        int
	Title     string
	Desc      string
	Date      string
	Timestamp int64
	Url       string
	ThumbUrl  string
}

type BBListArticleResponse struct {
	ArticleList []Article
}

type Company struct {
	Id       int
	Name     string
	Desc     string
	ThumbUrl string
}

type BBListCompanyRequest struct {
	StartIndex int
	PageSize   int
}

type BBListCompanyResponse struct {
	CompanyList []Company
}

type BBCreateCompanyResponse struct {
	Id       int64
	ThumbUrl string
}

type Insurance struct {
	Id        int
	NameZHCN  string
	NameEN    string
	Desc      string
	Company   string
	Timestamp int64
	ThumbUrl  string
	Type      string
}

type Type struct {
	Id   int64
	Name string
}

type BBListInsuranceRequest struct {
	StartIndex int
	PageSize   int
}

type BBListInsuranceResponse struct {
	InsuranceList []Insurance
}

type BBCreateInsuranceResponse struct {
	Id       int64
	ThumbUrl string
}

type Comment struct {
	LocalId   int64
	ServId    int64
	Uin       int64
	Content   string
	Score     int
	Timestamp int64
	ViewCount int
	Flags     int64
}

type BBCreateCommentRequest struct {
	Comment Comment
}

type BBCreateCommentResponse struct {
	ServId int64
}

type BBListCommentRequest struct {
	StartIndex int
	PageSize   int
}

type BBListCommentResponse struct {
	CommentList []Comment
}

type BBViewCommentRequest struct {
	ServId int64
}
