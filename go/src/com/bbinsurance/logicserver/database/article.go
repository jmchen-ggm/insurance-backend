package database

import (
	"com/bbinsurance/log"
	"container/list"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Article struct {
	Id       int
	Title    string
	Desc     string
	Url      string
	ThumbUrl string
}

func InsertArticle(article string, desc string, url string, thumbUrl string) (int64, error) {
	db, err := sql.Open("sqlite3", "./logic.db")
	defer db.Close()
	if err != nil {
		log.Error("Open DB Error %s", err)
		return -1, err
	}

	stmt, err := db.Prepare("INSERT INTO Article (Title, Desc, Url, ThumbUrl) VALUES (?, ?, ?, ?);")
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
	db, err := sql.Open("sqlite3", "./logic.db")
	defer db.Close()
	if err != nil {
		log.Error("Open DB Error %s", err)
		return
	}

	stmt, err := db.Prepare("UPDATE Article SET thumbUrl=? WHERE id=?;")
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return
	}

	result, err := stmt.Exec(thumbUrl, id)
	if err != nil {
		log.Error("Prepare Exec Error %s", err)
		return
	} else {
		log.Info("Result=%s", result)
	}
}

func GetAllArticle() *list.List {
	fmt.Println("GetAllArticle")
	db, err := sql.Open("sqlite3", "./logic.db")
	CheckErr(err)

	rows, err := db.Query("SELECT * FROM Article;")
	var articleList = list.New()

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Url, &article.ThumbUrl)
		CheckErr(err)
		articleList.PushBack(article)
	}
	rows.Close()
	return articleList
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
