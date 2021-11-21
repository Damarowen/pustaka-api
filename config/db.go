package config

import (
	"log"
	"pustaka-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConn struct {
	DbSQL  *gorm.DB

	// other db...
	// DbMasterFDBR *gorm.DB
	// DbSlaveFDBR  *gorm.DB
	// DbMasterForum *gorm.DB
	// DbSlaveForum  *gorm.DB
	// DbGiveaway *gorm.DB
	// DbAdmin    *gorm.DB

}



func ConnectDatabase() (data *DbConn, err error){

	dbSource := &DbConn{}

	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	dbSource.DbSQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Panic("Failed to connect to database!")
	}

	dbSource.DbSQL.AutoMigrate(&models.Book{})
	log.Println("CONECTED TO DB")


	return dbSource, err
}
