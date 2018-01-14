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
const CommentUpTableVersionKey = "CommentUpTableVersionKey"
const CommentReplyTableVersionKey = "CommentReplyTableVersionKey"

const CurrentArticleTableVersion = "0"
const CurrentCompanyTableVersion = "1"
const CurrentInsuranceTableVersion = "1"
const CurrentInsuranceTypeTableVersion = "2"
const CurrentCommentTableVersion = "2"
const CurrentCommentUpTableVersion = "0"
const CurrentCommentReplyTableVersion = "0"

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
	CheckTable(CommentUpTableName, CommentUpTableVersionKey, CurrentCommentUpTableVersion)
	CheckTable(CommentReplyTableName, CommentReplyTableVersionKey, CurrentCommentReplyTableVersion)

	var createArticleSql = "CREATE TABLE IF NOT EXISTS Article(Id INTEGER PRIMARY KEY AUTOINCREMENT, Title TEXT NOT NULL, Desc TEXT NOT NULL, Date TEXT NOT NULL, Timestamp INTEGER, Url TEXT NOT NULL, ThumbUrl TEXT NOT NULL);"
	_, err = db.Exec(createArticleSql, nil)
	if err != nil {
		log.Error("Create Article Error: sql = %s, err = %s", createArticleSql, err)
	} else {
		log.Info("Create Article Table Success sql = %s", createArticleSql)
	}

	var createCompanySql = "CREATE TABLE IF NOT EXISTS Company(Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT Not NULL, Desc TEXT NOT NULL, ThumbUrl TEXT NOT NULL, Flags INTEGER, DetailData TEXT);"
	_, err = db.Exec(createCompanySql, nil)
	if err != nil {
		log.Error("Create Company Error: sql = %s, err = %s", createCompanySql, err)
	} else {
		log.Info("Create Company Table Success sql = %s", createCompanySql)
	}
	var createInsuranceSql = "CREATE TABLE IF NOT EXISTS Insurance(Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT, Desc TEXT, InsuranceTypeId INTEGER, CompanyId INTEGER, AgeFrom INTEGER, AgeTo, INTEGER, AnnualCompensation INTEGER, AnnualPremium INTEGER,  Flags INTEGER, Timestamp INTEGER, ThumbUrl TEXT NOT NULL, DetailData TEXT NOT NULL);"
	_, err = db.Exec(createInsuranceSql, nil)
	if err != nil {
		log.Error("Create Insurance Error: sql = %s, err = %s", createInsuranceSql, err)
	} else {
		log.Info("Create Insurance Table Success sql = %s", createInsuranceSql)
	}
	var createInsuranceTypeSql = "CREATE TABLE IF NOT EXISTS InsuranceType(Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT Not NULL, Desc Text, ThumbUrl Text, Flags INTEGER, DetailData TEXT);"
	_, err = db.Exec(createInsuranceTypeSql, nil)
	if err != nil {
		log.Error("Create InsuranceType Error: sql = %s, err = %s", createInsuranceTypeSql, err)
	} else {
		log.Info("Create InsuranceType Table Success sql = %s", createInsuranceTypeSql)
	}

	// 评论
	var createCommentSql = "CREATE TABLE IF NOT EXISTS Comment(Id INTEGER PRIMARY KEY AUTOINCREMENT, Uin INTEGER, Content TEXT NOT NULL, CompanyId INTEGER, InsuranceTypeId INTEGER, TotalScore INTEGER, Score1 INTEGER, Score2 INTEGER, Score3 INTEGER, Score4 INTEGER, Timestamp INTEGER, UpCount INTEGER, ViewCount INTEGER, ReplyCount INTEGER, Flags INTEGER);"
	_, err = db.Exec(createCommentSql, nil)
	if err != nil {
		log.Error("Create Comment Error: sql = %s, err = %s", createCommentSql, err)
	} else {
		log.Info("Create Comment Table Success sql = %s", createCommentSql)
	}

	var createCommentTimestampIndex = "CREATE INDEX IF NOT EXISTS Comment_Timestamp ON Comment(Timestamp)"
	db.Exec(createCommentTimestampIndex, nil)

	var createCommentUpSql = "CREATE TABLE IF NOT EXISTS CommentUp(Id INTEGER PRIMARY KEY AUTOINCREMENT, Uin INTEGER, CommentId INTEGER, Timestamp INTEGER);"
	_, err = db.Exec(createCommentUpSql, nil)
	if err != nil {
		log.Error("Create CommentUp Table Error: sql = %s, err = %s", createCommentUpSql, err)
	} else {
		log.Info("Create CommentUp Table Success sql = %s", createCommentUpSql)
	}

	var createCommentUpUinCommentIdIndex = "CREATE INDEX IF NOT EXISTS CommentUp_Uin_CommentId ON CommentUp(Uin, CommentId)"
	db.Exec(createCommentUpUinCommentIdIndex, nil)

	var createCommentReplySql = "CREATE TABLE IF NOT EXISTS CommentReply(Id INTEGER PRIMARY KEY AUTOINCREMENT, Uin INTEGER, ReplyUin INTEGER, CommentId INTEGER, Content TEXT NOT NULL, TimeStamp INTEGER);"
	_, err = db.Exec(createCommentReplySql, nil)
	if err != nil {
		log.Error("Create CommentReply Error: sql = %s, err = %s", createCommentReplySql, err)
	} else {
		log.Info("Create CommentReply Table Success sql = %s", createCommentReplySql)
	}
	var createCommentReplyCommentIdIndex = "CREATE INDEX IF NOT EXISTS CommentReply_CommentId ON CommentReply(CommentId)"
	db.Exec(createCommentReplyCommentIdIndex, nil)
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
