package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/util"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const ArticleTableVersionKey = "ArticleTableVersionKey"
const CompanyTableVersionKey = "ComponyTableVersionKey"
const InsuranceTableVersionKey = "InsuranceTableVersionKey"
const InsuranceTypeTableVersionKey = "InsuranceTypeTableVersionKey"
const CommentTableVersionKey = "CommentTableVersionKey"

const CurrentArticleTableVersion = "0"
const CurrentCompanyTableVersion = "0"
const CurrentInsuranceTableVersion = "0"
const CurrentInsuranceTypeTableVersion = "1"
const CurrentCommentTableVersion = "1"

var db *sql.DB

func InitDB() {
	log.Info("InitDB Start")
	var err error
	db, err = sql.Open("sqlite3", "./logic.db")
	if err != nil {
		log.Error("open db err = %s", err)
	} else {
		log.Info("open db success")
	}

	var createConfigSql = "CREATE TABLE IF NOT EXISTS Config(Key TEXT PRIMARY KEY, Value Text NOT NULL)"
	_, err = db.Exec(createConfigSql, nil)
	if err != nil {
		log.Error("Create Config Error: sql = %s, err = %s", createConfigSql, err)
	} else {
		log.Info("Create Config Table Success sql = %s", createConfigSql)
	}

	CheckTable(ArticleTableName, ArticleTableVersionKey, CurrentArticleTableVersion)
	CheckTable(CompanyTableName, CompanyTableVersionKey, CurrentCompanyTableVersion)
	CheckTable(InsuranceTableName, InsuranceTableVersionKey, CurrentInsuranceTableVersion)
	CheckTable(InsuranceTypeTableName, InsuranceTypeTableVersionKey, CurrentInsuranceTypeTableVersion)
	CheckTable(CommentTableName, CommentTableVersionKey, CurrentCommentTableVersion)

	var createArticleSql = "CREATE TABLE IF NOT EXISTS Article(Id INTEGER PRIMARY KEY AUTOINCREMENT, Title TEXT NOT NULL, Desc TEXT NOT NULL, Date TEXT NOT NULL, Timestamp INTEGER, Url TEXT NOT NULL, ThumbUrl TEXT NOT NULL);"
	_, err = db.Exec(createArticleSql, nil)
	if err != nil {
		log.Error("Create Article Error: sql = %s, err = %s", createArticleSql, err)
	} else {
		log.Info("Create Article Table Success sql = %s", createArticleSql)
	}

	var createCompanySql = "CREATE TABLE IF NOT EXISTS Company(Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT Not NULL, Desc TEXT NOT NULL, ThumbUrl TEXT NOT NULL);"
	_, err = db.Exec(createCompanySql, nil)
	if err != nil {
		log.Error("Create Company Error: sql = %s, err = %s", createCompanySql, err)
	} else {
		log.Info("Create Company Table Success sql = %s", createCompanySql)
	}
	var createInsuranceSql = "CREATE TABLE IF NOT EXISTS Insurance(Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT Not NULL, Desc TEXT NOT NULL, InsuranceTypeId INTEGER, CompanyId INTEGER, Timestamp INTEGER, ThumbUrl TEXT NOT NULL, DetailData TEXT NOT NULL);"
	_, err = db.Exec(createInsuranceSql, nil)
	if err != nil {
		log.Error("Create Insurance Error: sql = %s, err = %s", createInsuranceSql, err)
	} else {
		log.Info("Create Insurance Table Success sql = %s", createInsuranceSql)
	}
	var createInsuranceTypeSql = "CREATE TABLE IF NOT EXISTS InsuranceType(Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT Not NULL, Desc Text, ThumbUrl Text);"
	_, err = db.Exec(createInsuranceTypeSql, nil)
	if err != nil {
		log.Error("Create InsuranceType Error: sql = %s, err = %s", createInsuranceTypeSql, err)
	} else {
		log.Info("Create InsuranceType Table Success sql = %s", createInsuranceTypeSql)
	}
	var createCommentSql = "CREATE TABLE IF NOT EXISTS Comment(Id INTEGER PRIMARY KEY AUTOINCREMENT, Uin INTEGER, Content TEXT NOT NULL, TotalScore INTEGER, Score1 INTEGER, Score2 INTEGER, Score3 INTEGER, Score4 INTEGER, Timestamp INTEGER, UpCount INTEGER, ViewCount INTEGER, ReplyCount INTEGER, Flags INTEGER);"
	_, err = db.Exec(createCommentSql, nil)
	if err != nil {
		log.Error("Create Comment Error: sql = %s, err = %s", createCommentSql, err)
	} else {
		log.Info("Create Comment Table Success sql = %s", createCommentSql)
	}

	var createCommentTimestampIndex = "CREATE INDEX IF NOT EXISTS Comment_Timestamp ON Comment(Timestamp)"
	db.Exec(createCommentTimestampIndex, nil)

	var createSubCommentSql = "CREATE TABLE IF NOT EXISTS SubComment(Id INTEGER PRIMARY KEY AUTOINCREMENT, Uin INTEGER, ReplyUin INTEGER, CommentId INTEGER, Content TEXT NOT NULL, TimeStamp INTEGER);"
	_, err = db.Exec(createSubCommentSql, nil)
	if err != nil {
		log.Error("Create SubComment Error: sql = %s, err = %s", createSubCommentSql, err)
	} else {
		log.Info("Create SubComment Table Success sql = %s", createSubCommentSql)
	}
}

func GetDB() *sql.DB {
	return db
}

func DropTable(tableName string) {
	sql := fmt.Sprintf("DROP TABLE %s;", tableName)
	db.Exec(sql, nil)
	log.Info("DropTable %s", sql)
}

func GetTableVersion(key string) string {
	version, err := GetConfig(key)
	if err != nil {
		version = "0"
	}
	if util.IsEmpty(version) {
		version = "0"
	}
	return version
}

func CheckTable(tableName string, key string, currentVersion string) {
	verison := GetTableVersion(key)
	if verison != currentVersion {
		DropTable(tableName)
	}
	SetConfig(key, currentVersion)
}
