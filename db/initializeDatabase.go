package db

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once

var StockTrackingDb *sql.DB

func ConnectDatabase() {
	once.Do(func() {
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/stockTrackingDb")
		if err != nil {
			panic(err.Error())
		}
		StockTrackingDb = db
	})

}
