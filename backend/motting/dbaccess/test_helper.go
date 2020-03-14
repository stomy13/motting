package dbaccess

import (
	"github.com/jinzhu/gorm"
)

func ConnectGormInTest() *gorm.DB {
	conargs := &ConnectArgs{
		Address:  "localhost",
		Port:     "33333",
		DBName:   "motting",
		User:     "motting",
		Password: "motting"}
	conargs.SetDefault()
	return ConnectGorm()
}
