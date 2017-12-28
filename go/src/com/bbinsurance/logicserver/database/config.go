package database

import (
	"com/bbinsurance/log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const ConfigTableName = "Config"

func SetConfig(key string, value string) error {
	sql := fmt.Sprintf("INSERT INTO %s (Key, Value) VALUES (?, ?);", ConfigTableName)
	stmt, err := GetDB().Prepare(sql)
	defer stmt.Close()
	if err != nil {
		log.Error("Prepare SQL Error %s", err)
		return err
	} else {
		_, err := stmt.Exec(key, value)
		if err != nil {
			log.Error("Prepare Exec Error %s", err)
			return err
		} else {
			return nil
		}
	}
}

func GetConfig(key string) (string, error) {
	sql := fmt.Sprintf("SELECT value FROM %s WHERE key = ?", ConfigTableName)
	log.Info("GetConfig key=%s", key)
	rows, err := GetDB().Query(sql, key)
	defer rows.Close()
	if err != nil {
		log.Error("GetConfig err %s", err)
		return "", err
	} else {
		if rows.Next() {
			var value string
			rows.Scan(&value)
			log.Info("GetConfig Key=%s value=%s", key, value)
			return value, nil
		} else {
			return "", nil
		}
	}
}
