package model

import (
	"testing"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
)

func TestPushTime(t *testing.T) {
	pt := &PushTime{
		UserID: "User1",
		PushAt: "10:10",
	}
	t.Log(pt.UserID)
	t.Log(pt.PushAt)
}

func TestPushTimeCreateTable(t *testing.T) {
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&PushTime{})
}
