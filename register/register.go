package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	gorm.Model
	Account  string
	Password string
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

// 自动创建数据库
func CreDb() {
	DB.AutoMigrate(&LoginUser{})
	DB.Create(&LoginUser{})
}

// 用户注册时新增
func AddUser(loginuser *LoginUser) {
	DB.Create(&loginuser)
	return
}

func UserDetailByName(account string) (loginuser LoginUser) {
	DB.Where("account = ?", account).First(&loginuser)
	return
}

func UserDetail(id uint) (loginuser LoginUser) {
	DB.Where("id = ?", id).First(&loginuser)
	return
}

func GetUserTotal() (int, error) {
	var count int
	if err := DB.Model(&LoginUser{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 密码加密
func Encrypt(source string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashPwd), err
}

// 密码比对 (传入未加密的密码即可)
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// gin session key
const KEY = "AEN233"

// 使用 Cookie 保存 session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("SAMPLE", store)
}

// session中间件 AuthSessionMiddle
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		if sessionValue == nil {
			c.JSON(200, "你无权访问,请登录后再访问,我是负责拦你的管事员")
			//c.String(200,"你无权访问,请登录后再访问!!")
			c.Abort()
			return
		}
		// 设置简单的变量
		c.Set("userId", sessionValue.(uint))

		c.Next()
		return
	}
}

// 注册和登陆时都需要保存seesion信息
func SaveAuthSession(c *gin.Context, id uint) {
	session := sessions.Default(c)
	session.Set("userId", id)
	session.Save()
}

// 退出时清除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("userId"); sessionValue == nil {
		return false
	}
	return true
}

func GetSessionUserId(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return 0
	}
	return sessionValue.(uint)
}

func Register(c *gin.Context) {
	var loginuser LoginUser
	loginuser.Account = c.PostForm("account")
	loginuser.Password = c.PostForm("password")

	if hasSession := HasSession(c); hasSession == true {
		c.JSON(200, "用户已登录")
		return
	}

	if existUser := UserDetailByName(loginuser.Account); existUser.ID != 0 {
		c.JSON(200, "用户已存在请重新注册")
		return
	}

	if c.PostForm("password") != c.PostForm("password") {
		c.JSON(200, "2次密码输入不一致")
		return
	}

	if pwd, err := Encrypt(c.Request.FormValue("password")); err == nil {
		loginuser.Password = pwd
	}

	AddUser(&loginuser)
	SaveAuthSession(c, loginuser.ID)
	c.JSON(200, "注册成功")
}

func Router() {
	r := gin.Default()

	r.POST("/register", Register)
	r.Run(":82")
}

func main() {
	DbInit()
	Router()
}
