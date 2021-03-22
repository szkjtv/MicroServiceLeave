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

	// CreDb()
	return DB

}

//删除函数
func Delete(c *gin.Context) {
	db := DbInit()
	id := c.Param("id")
	var delete Address
	db.Where(id).Delete(&delete)
	defer db.Close()
	c.JSON(200, "删除成功")
	// c.Redirect(http.StatusMovedPermanently, "/user/auth/address")
}

// 访问路由
func Router() {

	r := gin.Default()
	// r.LoadHTMLGlob("view/**/*")
	// r.Static("/static", "./static")

	// r.POST("/add", add)
	r.GET("/delete/:id", Delete)

	r.Run(":86")
}

func main() {
	Router()
	// for {
	// 	time.Sleep(0 * time.Microsecond)
	// }
}
