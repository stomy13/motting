package witticism

import (
	"github.com/MasatoTokuse/motting/motting/domain/user"
	"github.com/MasatoTokuse/motting/motting/domain/witticism"
	"github.com/MasatoTokuse/motting/motting/error_response"
)

func createWitticism(command *AddWitticismCommand) (*witticism.Witticism, error) {
	validateErrors := error_response.NewValidateErrors()
	tellerName, err := witticism.NewTellerName(command.TellerName)
	validateErrors.Append("tellerName", err)

	sentence, err := witticism.NewSentence(command.Sentence)
	validateErrors.Append("sentence", err)

	userId := user.UserId(command.OwnerId)

	witticism, err := witticism.NewWitticism(
		tellerName,
		sentence,
		&userId,
	)
	validateErrors.Append("witticism", err)
	if validateErrors.HasError() {
		return nil, validateErrors
	}
	return witticism, nil
}
