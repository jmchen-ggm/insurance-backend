package webcommon

const ResponseCodeSuccess = 0
const ResponseCodeServerError = -1
const ResponseCodeRequestInvalid = -2
const ResponseCodeInvalidURI = -3
const ResponseCodeInvalidFunId = -4

const FuncListArticle = 1
const FuncListCompany = 2
const FuncListInsurance = 3
const FuncListComment = 4
const FuncCreateComment = 5
const FuncViewComment = 6
const FuncListInsuranceType = 7
const FuncGetHomeData = 8
const FuncUpComment = 9
const FuncGetCompanyById = 10
const FuncGetInsuranceTypeById = 11
const FuncReplyComment = 12
const FuncGetListCommentReply = 13

const FuncRegisterUser = 101
const FuncLogin = 102
const FuncGetUser = 103
const FuncBatchGetUser = 104
const FuncGetAllUser = 105

const FuncCreateArticle = 10001
const FuncCreateCompany = 10002
const FuncCreateInsurance = 10003
const FuncCreateInsuranceType = 100004

const UriData = "data-bin"
const UriCreateData = "create-bin"
const UriUser = "user-bin"

const FileServer = "http://120.78.175.235:8081/"
