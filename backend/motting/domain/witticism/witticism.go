package witticism

import (
	"github.com/MasatoTokuse/motting/motting/domain/user"
)

// Witticism 名言。より良い人生をつくるための言葉
type Witticism struct {
	id         WitticismId
	tellerName TellerName
	sentence   Sentence
	owner      Owner
}

// WitticismId 名言のID。このドメインのID
type WitticismId string
// TellerName 名言の発案者名。誰が言ったか
type TellerName string
// Sentence 名言の内容。何を言ったか
type Sentence string
// Owner 名言を登録したユーザーのID
type Owner user.UserId
