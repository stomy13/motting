package witticism

import (
	"github.com/MasatoTokuse/motting/motting/domain/user"
	"github.com/MasatoTokuse/motting/motting/domain/witticism"
)

func createWitticism(command *AddWitticismCommand) (*witticism.Witticism, error) {
	tellerName, err := witticism.NewTellerName(command.TellerName)
	if err != nil {
		return nil, err
	}

	sentence, err := witticism.NewSentence(command.Sentence)
	if err != nil {
		return nil, err
	}
	userId := user.UserId(command.OwnerId)

	return witticism.NewWitticism(
		tellerName,
		sentence,
		&userId,
	)
}
