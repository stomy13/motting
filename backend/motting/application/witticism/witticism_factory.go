package witticism

import (
	"github.com/MasatoTokuse/motting/motting/domain/user"
	"github.com/MasatoTokuse/motting/motting/domain/witticism"
	"github.com/MasatoTokuse/motting/motting/error_response"
)

func createWitticism(command *AddWitticismCommand) (*witticism.Witticism, error) {
	errorResponse := error_response.NewError()
	tellerName, err := witticism.NewTellerName(command.TellerName)
	errorResponse.Append("tellerName", err)

	sentence, err := witticism.NewSentence(command.Sentence)
	errorResponse.Append("sentence", err)

	userId := user.UserId(command.OwnerId)

	witticism, err := witticism.NewWitticism(
		tellerName,
		sentence,
		&userId,
	)
	errorResponse.Append("witticism", err)
	return witticism, err
}
