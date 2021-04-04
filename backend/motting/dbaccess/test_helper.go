package dbaccess

import (
	"github.com/jinzhu/gorm"
)

func ConnectGormInTest() *gorm.DB {
	conargs := &ConnectArgs{
		Address:  "motting-db-test",
		Port:     "3306",
		DBName:   "motting",
		User:     "motting",
		Password: "motting"}
	conargs.SetDefault()
	return ConnectGorm()
}
