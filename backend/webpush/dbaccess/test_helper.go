package dbaccess

import (
	"github.com/jinzhu/gorm"
)

func ConnectGormInTest() *gorm.DB {
	conargs := &ConnectArgs{
		Address:  "localhost",
		Port:     "33333",
		DBName:   "webpush",
		User:     "webpush",
		Password: "webpush"}
	conargs.SetDefault()
	return ConnectGorm()
}
