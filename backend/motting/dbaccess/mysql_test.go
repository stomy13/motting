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
		Port:     "33306",
		DBName:   "motting",
		User:     "motting",
		Password: "motting"}

	db := ConnectGorm(conargs)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Test{})
}
