package mysql

import (
	"fmt"
	"testing"

	"github.com/MasatoTokuse/motting/motting/domain/user"
	"github.com/MasatoTokuse/motting/motting/domain/witticism"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_NewWitticismRepository(t *testing.T) {
	connectTemplate := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	defaultConnectString := fmt.Sprintf(connectTemplate, "motting", "motting", "motting-db-test", "3306", "motting")
	db, err := gorm.Open(mysql.Open(defaultConnectString), &gorm.Config{})
	assert.Nil(t, err)

	repository := NewWitticismRepository(db)

	tellerName, _ := witticism.NewTellerName("tellerName")
	sentence, _ := witticism.NewSentence("sentence")
	ownerId := user.UserId("ownerId")
	witticism, _ := witticism.NewWitticism(tellerName, sentence, &ownerId)

	repository.Save(witticism)
}
