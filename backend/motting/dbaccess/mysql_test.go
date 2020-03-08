package dbaccess

import (
	"testing"

	"github.com/jinzhu/gorm"
)

type Test struct {
	gorm.Model
	Test string `gorm:"size255"`
}

// 疎通確認
func Test_ConnectGorm_1(t *testing.T) {
	conargs := &ConnectArgs{
		Address:  "localhost",
		Port:     "33333",
		DBName:   "motting",
		User:     "motting",
		Password: "motting"}
	conargs.SetDefault()
	db := ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Test{})
}
