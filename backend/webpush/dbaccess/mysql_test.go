package dbaccess

import (
	"testing"

	"github.com/MasatoTokuse/motting/webpush/model"
)

func Test_CreateTable_Subscription(t *testing.T) {
	db := ConnectGormInTest()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&model.Subscription{})
}
