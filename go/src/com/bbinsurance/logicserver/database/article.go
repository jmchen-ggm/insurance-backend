package database

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/protocol"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func InsertArticle(article string, desc string, url string, thumbUrl string) (int64, error) {
	stmt, err := GetDB().Prepare("INSERT INTO Article (Title, Desc, Url, ThumbUrl) VALUES (?, ?, ?, ?);")
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return -1, err
	}

	result, err := stmt.Exec(article, desc, url, thumbUrl)
	if err != nil {
		log.Error("Prepare Exec Error %s", err)
		return -1, err
	}
	id, err := result.LastInsertId()
	return id, err
}

func UpdateArticleThumbUrl(id int64, thumbUrl string) {
	log.Info("UpdateArticleThumbUrl: id=%d thumbUrl=%s", id, thumbUrl)
	stmt, err := GetDB().Prepare("UPDATE Article SET thumbUrl=? WHERE id=?;")
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return
	}

	_, err = stmt.Exec(thumbUrl, id)
	if err != nil {
		log.Error("Prepare Exec Error %s", err)
		return
	} else {
		log.Info("UpdateArticleThumbUrl Success")
	}
}

func GetListArticle(startIndex int, length int) []protocol.Article {
	var sql string
	if (length == -1) {
		sql = fmt.Sprintf("SELECT * FROM Article")
	} else {
		sql = fmt.Sprintf("SELECT * FROM Article LIMIT %d OFFSET %d", length, startIndex)
	}
	log.Info("GetListArticle sql=%s", sql)
	rows, err := GetDB().Query(sql)
	if err != nil {
		log.Error("GetListArticle err %s", err)
	}
	var articleList []protocol.Article
	for rows.Next() {
		var article protocol.Article
		rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Url, &article.ThumbUrl)
		articleList = append(articleList, article)
	}
	rows.Close()
	return articleList
}
