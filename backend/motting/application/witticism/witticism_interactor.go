package witticism

import (
	"github.com/MasatoTokuse/motting/motting/domain/user"
	"github.com/MasatoTokuse/motting/motting/domain/witticism"
)

type WitticismUsecaseInteractor struct {
	WitticismRepository witticism.WitticismRepositoryInterface
}

func NewWitticismUsecaseInteractor(witticismRepository witticism.WitticismRepositoryInterface) WitticismUsecaseInterface {
	return &WitticismUsecaseInteractor{WitticismRepository: witticismRepository}
}

func (interractor *WitticismUsecaseInteractor) AddWitticism(command *AddWitticismCommand) error {
	tellerName, err := witticism.NewTellerName(command.TellerName)
	if err != nil {
		return err
	}
	sentence, err := witticism.NewSentence(command.Sentence)
	if err != nil {
		return err
	}
	userId := user.UserId(command.OwnerId)
	witticism, err := witticism.NewWitticism(
		tellerName,
		sentence,
		&userId,
	)
	if err != nil {
		return err
	}
	interractor.WitticismRepository.Save(witticism)
	return nil
}
