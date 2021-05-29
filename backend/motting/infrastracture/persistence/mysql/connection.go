package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection(hostName, port, dbName, userName, password string) *gorm.DB {
	mySqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", userName, password, hostName, port, dbName)

	connection, err := gorm.Open(mysql.Open(mySqlDSN), &gorm.Config{})
	panicIf(err)

	db, err := connection.DB()
	panicIf(err)

	err = db.Ping()
	panicIf(err)

	return connection
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
