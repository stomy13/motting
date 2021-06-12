package mysql

import (
	"github.com/MasatoTokuse/motting/motting/domain/user"
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

func (WitticismModel) TableName() string {
	return "witticisms"
}

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
	return NewDBErrorIfNotNil(result.Error)
}

func (witticismRepository *WitticismRepository) FindByOwnerId(ownerId *user.UserId) ([]*witticism.Witticism, error) {
	var models WitticismModels
	result := witticismRepository.db.Where("owner_id = ?", ownerId).Find(&models)
	if err := NewDBErrorIfNotNil(result.Error); err != nil {
		return nil, err
	}

	witticisms, err := models.toWitticisms()
	if err != nil {
		return nil, err
	}
	return witticisms, nil
}

func (model *WitticismModel) toWitticism() (*witticism.Witticism, error) {
	return witticism.NewWitticismWithUUID(model.ID, model.TellerName, model.Sentence, model.OwnerId)
}

type WitticismModels []WitticismModel

func (models *WitticismModels) toWitticisms() ([]*witticism.Witticism, error) {
	var witticisms []*witticism.Witticism
	for _, model := range *models {
		witticism, err := model.toWitticism()
		if err != nil {
			return nil, err
		}
		witticisms = append(witticisms, witticism)
	}
	return witticisms, nil
}
