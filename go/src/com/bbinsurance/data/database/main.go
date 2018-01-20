package main

import (
	"com/bbinsurance/data/constants"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/util"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 初始化变量
	constants.InitConstants()

	fmt.Printf("%s %s\n", constants.STATIC_FOLDER, constants.LOGIC_DB_PATH)

	HandleInsuranceType()
	HandleAritcle()
}

func HandleAritcle() {
	articleJsonStr, _ := util.FileGetContent(constants.STATIC_FOLDER + "/data/articles.json")
	fmt.Printf("%s\n", articleJsonStr)
	var listArticleResponse protocol.BBListArticleResponse
	json.Unmarshal(util.StringToBytes(articleJsonStr), &listArticleResponse)
	articleList := listArticleResponse.ArticleList
	for i := 0; i < len(articleList); i++ {
		article, _ := InsertArticle(articleList[i])
		if article.Id != -1 {
			fmt.Printf("Insert Success %s\n", util.ObjToString(article))
		}
	}
}

func HandleInsuranceType() {
	insuranceTypeJsonStr, _ := util.FileGetContent(constants.STATIC_FOLDER + "/data/insurance-types.json")
	var listInsuranceTypeResponse protocol.BBListInsuranceTypeResponse
	json.Unmarshal(util.StringToBytes(insuranceTypeJsonStr), &listInsuranceTypeResponse)
	insuranceTypeList := listInsuranceTypeResponse.InsuranceTypeList
	for i := 0; i < len(insuranceTypeList); i++ {
		insuranceType, _ := InsertInsuranceType(insuranceTypeList[i])
		if insuranceType.Id != -1 {
			fmt.Printf("Insert Success %s\n", util.ObjToString(insuranceType))
		}
	}
}

func InsertArticle(article protocol.Article) (protocol.Article, error) {
	db, _ := sql.Open("sqlite3", constants.LOGIC_DB_PATH)
	sql := fmt.Sprintf("INSERT OR REPLACE INTO Article (Id, Title, Desc, Date, Timestamp, Url, ThumbUrl, ViewCount) VALUES (?, ?, ?, ?, ?, ?, ?);")
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		article.Id = -1
		return article, err
	} else {
		result, err := stmt.Exec(article.Id, article.Title, article.Desc, article.Date, article.Timestamp, article.Url, article.ThumbUrl, article.ViewCount)
		if err != nil {
			article.Id = -1
			return article, err
		} else {
			article.Id, err = result.LastInsertId()
			return article, err
		}
	}
}

func InsertInsuranceType(insuranceType protocol.InsuranceType) (protocol.InsuranceType, error) {
	db, _ := sql.Open("sqlite3", constants.LOGIC_DB_PATH)
	sql := fmt.Sprintf("INSERT OR REPLACE INTO InsuranceType (Id, Name, Desc, ThumbUrl, Flags, DetailData) VALUES (?, ?, ?, ?, ?, ?);")
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		insuranceType.Id = -1
	} else {
		result, err := stmt.Exec(insuranceType.Id, insuranceType.Name, insuranceType.Desc, insuranceType.ThumbUrl, insuranceType.Flags, insuranceType.DetailData)
		if err != nil {
			insuranceType.Id = -1
		} else {
			insuranceType.Id, err = result.LastInsertId()
		}
	}
	return insuranceType, err
}
