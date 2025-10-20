package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// func Connect() {
// 	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?")
// 	if err != nil {
// 		panic(err)

// 	}
// 	db = d
// }

func Connect() {
	dsn := "root:EO0549200629.m@tcp(127.0.0.1:3307)/simplerest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
