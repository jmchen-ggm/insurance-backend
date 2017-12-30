package protocol

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
	Id       int64
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
	Company Company
}

type Insurance struct {
	Id                int64
	Name              string
	Desc              string
	InsuranceTypeId   int64
	InsuranceTypeName string
	CompanyId         int64
	CompanyName       string
	Timestamp         int64
	ThumbUrl          string
	DetailData        string
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
	Insurance Insurance
}

type Comment struct {
	Id         int64
	Uin        int64
	Content    string
	TotalScore int
	Score1     int
	Score2     int
	Score3     int
	Timestamp  int64
	UpCount    int
	ViewCount  int
	ReplyCount int
	Flags      int64
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
	ServerId int64
}

type InsuranceType struct {
	Id       int64
	Name     string
	Desc     string
	ThumbUrl string
}

type BBCreateInsuranceTypeResponse struct {
	InsuranceType InsuranceType
}

type BBListInsuranceTypeReponse struct {
	InsuranceTypeList []InsuranceType
}
