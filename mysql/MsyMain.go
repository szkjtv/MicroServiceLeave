package main

import (
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func DbInit() (db *gorm.DB) {
	DB, err = gorm.Open("mysql", "root:.aA1451418@tcp(123.207.88.76:3306)/updateflow?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()
	// fmt.Println("运行数据库")
	return DB

}

func main() {
	DbInit()
}
