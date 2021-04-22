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

//显示界面
func Showaddhtml(c *gin.Context) {
	c.HTML(200, "ShowQuery", nil)

}

//s查询所有数据
func ShowQuery(c *gin.Context) {
	db := DbInit()
	var queryaddress []Address
	db.Find(&queryaddress)
	// c.HTML(200, "ShowQuery", queryaddress)
	c.JSON(200, queryaddress)
	defer db.Close()
}

//查询单条数据
func QueryOne(c *gin.Context) {
	db := DbInit()
	var see Address
	id := c.Param("id")
	db.Find(&see, id)
	// c.HTML(200, "see.htm", see)
	c.JSON(200, see)
	defer db.Close()
}

// 访问路由
func Router() {

	r := gin.Default()
	// r.LoadHTMLGlob("view/**/*")
	// r.LoadHTMLFiles("veiw/showaddress.html")
	// r.Static("/static", "./static")

	r.GET("/Showadd", Showaddhtml)
	r.GET("/ShowQuery", ShowQuery) //查询所有数据方法成功可靠
	// http://127.0.0.1:88/QueryOne/id=15  查询单条数据接口方式
	r.GET("/QueryOne/:id", QueryOne) //查询单一数据方法成功可行

	r.Run(":88")
}

func main() {
	Router()
}
