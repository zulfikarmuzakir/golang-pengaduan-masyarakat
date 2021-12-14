package database

import (
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/db_pengaduan"), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Pengaduan{}, &models.Tanggapan{}, &models.Image{})
}
