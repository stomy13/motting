package dbaccess

import (
	"github.com/jinzhu/gorm"
)

func ConnectGormInTest() *gorm.DB {
	conargs := &ConnectArgs{
		Address:  "motting-db-test",
		Port:     "3306",
		DBName:   "webpush",
		User:     "webpush",
		Password: "webpush"}
	conargs.SetDefault()
	return ConnectGorm()
}
