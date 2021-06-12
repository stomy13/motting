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
	ownerId := user.UserId(command.OwnerId)
	witticism, err := witticism.NewWitticism(command.TellerName, command.Sentence, &ownerId)

	if err != nil {
		return err
	}
	return interractor.WitticismRepository.Save(witticism)
}
