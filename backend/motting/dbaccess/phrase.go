package dbaccess

import (
	"github.com/MasatoTokuse/motting/motting/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ParamPhrase struct {
	ID     string
	UserID string
	Text   string
	Author string
}

func QueryPhrases(db *gorm.DB, param *ParamPhrase) *[]model.Phrase {

	if param.ID != "" {
		db = db.Where("id = ?", param.ID)
	}
	if param.UserID != "" {
		db = db.Where("user_id = ?", param.UserID)
	}
	if param.Text != "" {
		db = db.Where("text LIKE ?", "%"+param.Text+"%")
	}
	if param.Author != "" {
		db = db.Where("author LIKE ?", "%"+param.Author+"%")
	}

	phrases := &[]model.Phrase{}
	db.Find(phrases)
	return phrases
}
