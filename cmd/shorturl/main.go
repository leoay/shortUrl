package main

import (
	urldata "ShortUrl/internal/dao"
	router "ShortUrl/internal/server"
)

func main() {
	router.InitRouter()
	urldata.DB_AutoMigrate()
}
