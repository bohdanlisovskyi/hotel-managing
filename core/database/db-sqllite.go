package database

import (
	"sync"

	"database/sql"

	"github.com/bohdanlisovskyi/hotel-managing/core/loger"
	_ "github.com/mattn/go-sqlite3"
)

var (
	once sync.Once
	db   *sql.DB
)

// connect to sqlite storage
func GetStorage() (*sql.DB, error) {
	var err error
	once.Do(func() {
		db, err = sql.Open("sqlite3", "core/database/hotel.db")
		if err != nil {
			loger.Log.Panicf("ERROR connection to SqlLite3 %s", err.Error())
		}
	})
	return db, err
}
