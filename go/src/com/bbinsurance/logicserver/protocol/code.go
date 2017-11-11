package protocol

const ResponseCodeSuccess = 0
const ResponseCodeServerError = -1
const ResponseCodeRequestInvalid = -2
const ResponseCodeInvalidURI = -3
const ResponseCodeInvalidFunId = -4

const FuncListArticle = 1
const FuncListCompany = 2
const FuncListInsurance = 3

const FuncCreateArticle = 101
const FuncCreateCompany = 102
const FuncCreateInsurance = 103

const UriArticleList = "data-bin"
const UriCreateData = "create-bin"
