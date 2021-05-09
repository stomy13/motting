package mysql

import (
	"github.com/MasatoTokuse/motting/motting/domain/witticism"
	"github.com/MasatoTokuse/motting/motting/infrastracture/persistence/abstract"
	"gorm.io/gorm"
)

type WitticismModel struct {
	abstract.Model
	TellerName string
	Sentence   string
	OwnerId    string
}

// TODO: created_at に値が入力されないなら自前で設定する
func createWitticismModel(witticism *witticism.Witticism) *WitticismModel {
	model := WitticismModel{
		TellerName: string(*witticism.TellerName),
		Sentence:   string(*witticism.Sentence),
		OwnerId:    string(*witticism.OwnerId),
	}
	model.ID = string(*witticism.Id)
	return &model
}

type WitticismRepository struct {
	db *gorm.DB
}

func NewWitticismRepository(db *gorm.DB) witticism.WitticismRepositoryInterface {
	return &WitticismRepository{db: db}
}

func (witticismRepository *WitticismRepository) Save(witticism *witticism.Witticism) error {
	model := createWitticismModel(witticism)
	result := witticismRepository.db.Create(model)
	return result.Error
}
