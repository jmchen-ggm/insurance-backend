package protocol

const ResponseCodeSuccess = 0
const ResponseCodeServerError = -1
const ResponseCodeRequestInvalid = -2
const ResponseCodeInvalidURI = -3
const ResponseCodeInvalidFunId = -4

const FuncListArticle = 1
const FuncListCompany = 2
const FuncCreateArticle = 3
const FuncCreateCompany = 4

const UriArticleList = "data-bin"
const UriCreateData = "create-bin"
