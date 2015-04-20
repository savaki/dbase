package dbase

import (
	"os"

	"github.com/savaki/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Default *gorm.DB

func init() {
	dialect := os.Getenv("DATABASE_DRIVER")
	if dialect == "" {
		dialect = "mysql"
	}

	open := os.Getenv("DATABASE_URL")
	if open == "" {
		log.Fatalln("DATABASE_URL not set!")
	}

	db, err := gorm.Open(dialect, open)
	if err != nil {
		log.Fatalln(err.Error())
	}

	Default = &db

	Default.DB().SetMaxIdleConns(10)
	Default.DB().SetMaxOpenConns(100)
}

func WithRollback(f func(db *gorm.DB) error) error {
	tx := Default.Begin()
	defer func() {
		tx.Rollback()
	}()
	return f(tx)
}
