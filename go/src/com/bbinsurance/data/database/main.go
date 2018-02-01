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

	// HandleInsuranceType()
	// HandleAritcle()
	// HandleCompany()
	HandleInsurance()
}

func HandleAritcle() {
	articleJsonStr, _ := util.FileGetContent(constants.STATIC_FOLDER + "/data/articles.json")
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

func HandleCompany() {
	companyJsonStr, _ := util.FileGetContent(constants.STATIC_FOLDER + "/data/companys.json")
	var listCompanyResponse protocol.BBListCompanyResponse
	json.Unmarshal(util.StringToBytes(companyJsonStr), &listCompanyResponse)
	companyList := listCompanyResponse.CompanyList
	for i := 0; i < len(companyList); i++ {
		company, _ := InsertCompany(companyList[i])
		if company.Id != -1 {
			fmt.Printf("Insert Success %s\n", util.ObjToString(company))
		}
	}
}

func HandleInsurance() {
	insuranceJsonStr, _ := util.FileGetContent(constants.STATIC_FOLDER + "/data/insurances.json")
	var listInsuranceResponse protocol.BBListInsuranceResponse
	json.Unmarshal(util.StringToBytes(insuranceJsonStr), &listInsuranceResponse)
	insuranceList := listInsuranceResponse.InsuranceList
	for i := 0; i < len(insuranceList); i++ {
		insurance, _ := InsertInsurance(insuranceList[i])
		if insurance.Id != -1 {
			fmt.Printf("Insert Success %s\n", util.ObjToString(insurance))
		}
	}
}

func InsertArticle(article protocol.Article) (protocol.Article, error) {
	db, _ := sql.Open("sqlite3", constants.LOGIC_DB_PATH)
	defer db.Close()
	sql := fmt.Sprintf("INSERT OR REPLACE INTO Article (Id, Title, Desc, Date, Timestamp, Url, ThumbUrl, ViewCount) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")
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
	defer db.Close()
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

func InsertCompany(company protocol.Company) (protocol.Company, error) {
	db, _ := sql.Open("sqlite3", constants.LOGIC_DB_PATH)
	defer db.Close()
	sql := fmt.Sprintf("INSERT OR REPLACE INTO Company (Id, Name, Desc, ThumbUrl, Flags, DetailData) VALUES (?, ?, ?, ?, ?, ?);")
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		company.Id = -1
	} else {
		result, err := stmt.Exec(company.Id, company.Name, company.Desc, company.ThumbUrl, company.Flags, company.DetailData)
		if err != nil {
			company.Id = -1
		} else {
			company.Id, err = result.LastInsertId()
		}
	}
	return company, err
}

func InsertInsurance(insurance protocol.Insurance) (protocol.Insurance, error) {
	db, _ := sql.Open("sqlite3", constants.LOGIC_DB_PATH)
	defer db.Close()
	sql := fmt.Sprintf("INSERT OR REPLACE INTO Insurance (Id, Name, Desc, InsuranceTypeId, CompanyId, AgeFrom, AgeTo, AnnualCompensation, AnnualPremium, Flags, Timestamp, ThumbUrl, DetailData) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		insurance.Id = -1
	} else {
		result, err := stmt.Exec(insurance.Id, insurance.Name, insurance.Desc, insurance.InsuranceTypeId, insurance.CompanyId,
			insurance.AgeFrom, insurance.AgeTo, insurance.AnnualCompensation, insurance.AnnualPremium,
			insurance.Flags, insurance.Timestamp, insurance.ThumbUrl, insurance.DetailData)
		if err != nil {
			insurance.Id = -1
		} else {
			insurance.Id, err = result.LastInsertId()
		}
	}
	return insurance, err
}
