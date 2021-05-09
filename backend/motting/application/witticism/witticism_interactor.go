package witticism

import (
	"github.com/MasatoTokuse/motting/motting/domain/witticism"
)

type WitticismUsecaseInteractor struct {
	WitticismRepository witticism.WitticismRepositoryInterface
}

func NewWitticismUsecaseInteractor(witticismRepository witticism.WitticismRepositoryInterface) WitticismUsecaseInterface {
	return &WitticismUsecaseInteractor{WitticismRepository: witticismRepository}
}

func (interractor *WitticismUsecaseInteractor) AddWitticism(command *AddWitticismCommand) error {
	witticism, err := createWitticism(command)
	if err != nil {
		return err
	}
	interractor.WitticismRepository.Save(witticism)
	return nil
}
