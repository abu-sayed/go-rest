package db

import _ "github.com/go-sql-driver/mysql" // database driver
import (
	"fmt"
	"database/sql"
	"sync"
	"log"
	"os"
)

var (
	db    *sql.DB
	once sync.Once
)

func GetDb() *sql.DB {

	once.Do(func() {
		dbSourceName := fmt.Sprintf("%s:%s@%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
		var err error
		db, err = sql.Open("mysql", dbSourceName)

		if err != nil {
			log.Println(err)
		}
	})

	return db
}