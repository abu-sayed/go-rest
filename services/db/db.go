package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"sync"
	"log"
)

var (
	db    *sql.DB
	once sync.Once
)

func GetDb() *sql.DB {

	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:123@/todos")

		if err != nil {
			log.Println(err)
		}
	})

	return db
}