package urldata

import (
	"ShortUrl/configs"
	models "ShortUrl/internal/model"
	"fmt"

	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func initdb() {
	mysqlserver := configs.GetConfig().DB.Server
	mysqluser := configs.GetConfig().DB.UserName
	mysqlpass := configs.GetConfig().DB.Passwd
	mysqlport := configs.GetConfig().DB.Ports
	mysqldb := configs.GetConfig().DB.DBName

	dsn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlserver + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		beego.Error(err)
		return
	}
}

func closeDB() {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB.Close()
}

func DB_AutoMigrate() {
	initdb()
	db.AutoMigrate(&models.UrlMap{})
	defer closeDB()
}

func StoreUrl(shorturl string, longurl string) {
	initdb()
	urlmaps := models.UrlMap{ShortUrl: shorturl, LongUrl: longurl}
	db.Create(&urlmaps)
	defer closeDB()
}

func ReadLongUrl(shorturl string) string {
	initdb()
	urlmaps := models.UrlMap{}
	db.Where("short_url=?", shorturl).Find(&urlmaps)
	defer closeDB()
	return urlmaps.LongUrl
}

func ReadShortUrl(longurl string) string {
	initdb()
	urlmaps := models.UrlMap{ShortUrl: "23233", LongUrl: "sdsdsd"}
	db.Where("long_url=?", longurl).Find(&urlmaps)
	defer closeDB()
	return urlmaps.ShortUrl
}
