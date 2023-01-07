package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


// Conection to database in MySQL/MariaDB
func Connection() {
	var err error
	dsn := fmt.Sprintf("user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil  {
		panic(err)
	}
}