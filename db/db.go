package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	db, err := gorm.Open("mysql", "{user}:{pwd}@tcp(host:3306)/{schema}?charset=utf8&parseTime=True") // ?charset=utf8&parseTime=True&loc=Local
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
}
