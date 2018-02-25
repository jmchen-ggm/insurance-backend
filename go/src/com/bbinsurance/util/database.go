package util

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const ConfigTableName = "Config"

func CreateDBConfigTable(db *sql.DB) {
	var createConfigSql = fmt.Sprintf("CREATE TABLE IF NOT EXISTS Config(Key TEXT PRIMARY KEY, Value Text NOT NULL)", ConfigTableName)
	db.Exec(createConfigSql, nil)
}

func SetDBConfig(db *sql.DB, key string, value string) error {
	sql := fmt.Sprintf("INSERT OR REPLACE INTO %s (Key, Value) VALUES (?, ?);", ConfigTableName)
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	} else {
		_, err := stmt.Exec(key, value)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}

func GetDBConfig(db *sql.DB, key string) (string, error) {
	sql := fmt.Sprintf("SELECT value FROM %s WHERE key = ?", ConfigTableName)
	rows, err := db.Query(sql, key)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return "", err
	} else {
		if rows.Next() {
			var value string
			rows.Scan(&value)
			return value, nil
		} else {
			return "", nil
		}
	}
}

func DropDBTable(db *sql.DB, tableName string) {
	sql := fmt.Sprintf("DROP TABLE %s;", tableName)
	db.Exec(sql, nil)
}

func GetDBTableVersion(db *sql.DB, key string) string {
	version, err := GetDBConfig(db, key)
	if err != nil {
		version = "0"
	}
	if IsEmpty(version) {
		version = "0"
	}
	return version
}

func CheckDBTable(db *sql.DB, tableName string, key string, currentVersion string) {
	verison := GetDBTableVersion(db, key)
	if verison != currentVersion {
		DropDBTable(db, tableName)
	}
	SetDBConfig(db, key, currentVersion)
}

func SetSequenceStartId(db *sql.DB, tableName string, startId int64) {
	sql := fmt.Sprintf("UPDATE sqlite_sequence SET seq = %d WHERE name = ?", startId)
	db.Exec(sql, tableName)
}
