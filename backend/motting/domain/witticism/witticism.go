package witticism

import (
	"github.com/MasatoTokuse/motting/motting/domain/user"
)

type Witticism struct {
	id         WitticismId
	tellerName TellerName
	sentence   Sentence
	owner      Owner
}

type WitticismId string
type TellerName string
type Sentence string
type Owner user.UserId
