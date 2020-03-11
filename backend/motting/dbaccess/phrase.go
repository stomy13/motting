package dbaccess

import (
	"fmt"

	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// PhrasesAll is return all phrases
func PhrasesAll(db *gorm.DB) *[]model.Phrase {
	phrases := &[]model.Phrase{}
	db.Find(phrases)
	fmt.Println(*phrases)
	return phrases
}
