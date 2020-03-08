package model

import (
	"testing"

	"github.com/MasatoTokuse/motting/motting/dbaccess"
)

func TestPhrase(t *testing.T) {
	phrase := &Phrase{
		UserID: "",
		Text:   "",
		Author: "",
	}
	t.Log(phrase.UserID)
	t.Log(phrase.Text)
	t.Log(phrase.Author)
}

func TestPhraseCreateTable(t *testing.T) {
	db := dbaccess.ConnectGormInTest()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Phrase{})
}
