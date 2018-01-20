package protocol

type BBCreateArticleResponse struct {
	Article Article
}

type BBListArticleRequest struct {
	StartIndex int
	PageSize   int
}

type Article struct {
	Id        int64
	Title     string
	Desc      string
	Date      string
	Timestamp int64
	Url       string
	ThumbUrl  string
	ViewCount int
}

type BBListArticleResponse struct {
	ArticleList []Article
}

type Company struct {
	Id         int64
	Name       string
	Desc       string
	ThumbUrl   string
	Flags      int64
	DetailData string
}

type BBListCompanyRequest struct {
	StartIndex int
	PageSize   int
}

type BBListCompanyResponse struct {
	CompanyList []Company
}

type BBCreateCompanyResponse struct {
	Company Company
}

type Insurance struct {
	Id                 int64
	Name               string
	Desc               string
	InsuranceTypeId    int64
	InsuranceTypeName  string
	CompanyId          int64
	CompanyName        string
	AgeFrom            int
	AgeTo              int
	AnnualCompensation int
	AnnualPremium      int
	Flags              int64
	Timestamp          int64
	ThumbUrl           string
	DetailData         string
}

type BBListInsuranceRequest struct {
	StartIndex int
	PageSize   int
}

type BBListInsuranceResponse struct {
	InsuranceList []Insurance
}

type BBCreateInsuranceResponse struct {
	Insurance Insurance
}

type Comment struct {
	Id                int64
	Uin               int64
	Content           string
	CompanyId         int64
	CompanyName       string
	InsuranceTypeId   int64
	InsuranceTypeName string
	TotalScore        int
	Score1            int
	Score2            int
	Score3            int
	Score4            int
	Timestamp         int64
	UpCount           int
	ViewCount         int
	ReplyCount        int
	Flags             int64
	IsUp              bool
}

type BBCreateCommentRequest struct {
	Comment Comment
}

type BBCreateCommentResponse struct {
	Comment Comment
}

type BBListCommentRequest struct {
	StartIndex int
	PageSize   int
}

type BBListCommentResponse struct {
	CommentList []Comment
}

type BBViewCommentRequest struct {
	Id int64
}

type BBViewCommentResponse struct {
	Comment Comment
}

type BBUpCommentRequest struct {
	CommentUp CommentUp
	IsUp      bool
}

type BBUpCommentResponse struct {
	Comment Comment
}

type BBReplyCommentRequest struct {
	CommentReply CommentReply
}

type BBReplyCommentResponse struct {
	Comment Comment
}

type InsuranceType struct {
	Id         int64
	Name       string
	Desc       string
	ThumbUrl   string
	Flags      int64
	DetailData string
}

type BBCreateInsuranceTypeResponse struct {
	InsuranceType InsuranceType
}

type BBListInsuranceTypeRequest struct {
	StartIndex int
	PageSize   int
}

type BBListInsuranceTypeResponse struct {
	InsuranceTypeList []InsuranceType
}

type BBGetHomeDataRequest struct {
}

type BBGetHomeDataResponse struct {
	BannerList           []Insurance
	TopCommentList       []Comment
	TopInsuranceTypeList []InsuranceType
	TopCompanyList       []Company
}

type BBGetCompanyRequest struct {
	Id int64
}

type BBGetCompanyResponse struct {
	Company Company
}

type BBGetInsuranceTypeRequest struct {
	Id int64
}

type BBGetInsuranceTypeResponse struct {
	InsuranceType InsuranceType
}

type CommentUp struct {
	Id        int64
	Uin       int64
	CommentId int64
	Timestamp int64
}

type CommentReply struct {
	Id        int64
	Uin       int64
	ReplyUin  int64
	CommentId int64
	Content   string
	Timestamp int64
}

type BBGetCommentReplyListRequest struct {
	CommentId int64
}

type BBGetCommentReplyListResponse struct {
	CommentReplyList []CommentReply
}
