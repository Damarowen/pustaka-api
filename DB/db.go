package DB

import (
	"log"
	"pustaka-api/book"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnDb() ( err error) {


	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Err connection admin", err.Error())
		return err
	}

	db.AutoMigrate(&book.Book{})

	log.Println("CONECTED TO DB")
	return nil
}
