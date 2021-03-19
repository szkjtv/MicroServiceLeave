package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

type Address struct {
	Id      int
	Name    string `json:"name"`
	Number  string `json:"number"`
	Address string `json:"address"`
	// Bz      string `json:"bz"`
	// CreatedAt time.Time
}

var DB *gorm.DB
var err error

// 链接数据库
func DbInit() (db *gorm.DB) {
	DB, err = gorm.Open("mysql", "root:aa1451418@tcp(127.0.0.1:3306)/updateflow?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	CreDb()
	return DB

}

// 自动仓库数据库
func CreDb() {

	DB.AutoMigrate(&Address{})

}

// 增加数据内容
func add(c *gin.Context) {

	db := DbInit()
	name := c.PostForm("name")
	number := c.PostForm("number")
	address := c.PostForm("address")
	newAdd := Address{Name: name, Number: number, Address: address}
	db.Create(&newAdd)
	c.JSON(200, "播入成功")
	defer db.Close()

}

// 访问路由
func Router() {

	r := gin.Default()
	// r.LoadHTMLGlob("view/**/*")
	// r.Static("/static", "./static")

	r.POST("/add", add)

	r.Run(":85")
}

func main() {
	Router()
	// for {
	// 	time.Sleep(0 * time.Microsecond)
	// }
}
