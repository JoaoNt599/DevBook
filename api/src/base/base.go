package base

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

func Cannect() (*sql.DB, error) {
	db, er := sql.Open("mysql", config.StringConnMysql)
	if er != nil {
		return nil, er
	}

	if er = db.Ping(); er != nil {
		db.Close()
		return nil, er
	}

	return db, nil
}
