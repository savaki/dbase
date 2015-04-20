package dbase

import (
	"os"

	"github.com/savaki/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	mysqlAccounts string
	mysqlServices string
)

var DB gorm.DB

func init() {
	dialect := os.Getenv("DATABASE_DRIVER")
	if dialect == "" {
		dialect = "mysql"
	}

	open := os.Getenv("DATABASE_URL")
	if open == "" {
		log.Fatalln("DATABASE_URL not set!")
	}

	var err error
	DB, err = gorm.Open(dialect, open)
	if err != nil {
		log.Fatalln(err.Error())
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}
