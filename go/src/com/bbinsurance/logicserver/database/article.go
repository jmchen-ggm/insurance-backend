package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	"com/bbinsurance/time"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const ArticleTableName = "Article"

func InsertArticle(title string, desc string, date string, url string, thumbUrl string) (int64, error) {
	sql := fmt.Sprintf("INSERT INTO %s (Title, Desc, Date, Timestamp, Url, ThumbUrl) VALUES (?, ?, ?, ?, ?, ?);", ArticleTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return -1, err
	} else {
		timestamp := time.GetTimestampInMilli()
		result, err := stmt.Exec(title, desc, date, timestamp, url, thumbUrl)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return -1, err
		} else {
			id, err := result.LastInsertId()
			return id, err
		}
	}
}

func UpdateArticleThumbUrl(id int64, thumbUrl string) {
	log.Info("UpdateArticleThumbUrl: id=%d thumbUrl=%s", id, thumbUrl)
	sql := fmt.Sprintf("UPDATE %s SET thumbUrl=? WHERE id= ?;", ArticleTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
	} else {
		_, err = stmt.Exec(thumbUrl, id)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
		} else {
			log.Info("UpdateArticleThumbUrl Success")
		}
	}
}

func GetListArticle(startIndex int, length int) []protocol.Article {
	var sql string
	if length == -1 {
		sql = fmt.Sprintf("SELECT * FROM %s", ArticleTableName)
	} else {
		sql = fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", ArticleTableName, length, startIndex)
	}
	log.Info("GetListArticle sql=%s", sql)
	rows, err := GetDB().Query(sql)
	defer rows.Close()
	var articleList []protocol.Article
	if err != nil {
		log.Error("GetListArticle err %s", err)
	} else {
		for rows.Next() {
			var article protocol.Article
			rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Date, &article.Timestamp, &article.Url, &article.ThumbUrl)
			articleList = append(articleList, article)
		}
		log.Info("GetListArticle %d ", len(articleList))
	}
	return articleList
}

func DeleteArticleById(id int64) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id=?", ArticleTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
	} else {
		_, err = stmt.Exec(id)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return
		} else {
			log.Info("RemoveArticleById %d Success", id)
		}
	}
}
