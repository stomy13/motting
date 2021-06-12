package test_helper

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

func GetConnection() *gorm.DB {
	if connection != nil {
		return connection
	}

	hostName := os.Getenv("TEST_DB_HOST")
	mySqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "motting", "motting", hostName, "3306", "motting")

	var err error
	connection, err = gorm.Open(mysql.Open(mySqlDSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return connection
}
