package witticism

import (
	"github.com/MasatoTokuse/motting/motting/domain/user"
	"github.com/google/uuid"
)

// Witticism 名言。より良い人生をつくるための言葉
type Witticism struct {
	id         *WitticismId
	tellerName *TellerName
	sentence   *Sentence
	ownerId    *user.UserId
}

func NewWitticism(
	tellerName TellerName,
	sentence Sentence,
	ownerId user.UserId,
) *Witticism {
	witticismId, _ := NewWitticismId()
	return &Witticism{
		id:         witticismId,
		tellerName: &tellerName,
		sentence:   &sentence,
		ownerId:    &ownerId,
	}
}

// WitticismId 名言のID。このドメインのID
type WitticismId string

func NewWitticismId() (*WitticismId, error) {
	witticism := WitticismId(uuid.NewString())
	return &witticism, witticism.valid()
}

// uuidであることの確認
func (witticismId *WitticismId) valid() error {
	_, err := uuid.Parse(string(*witticismId))
	return err
}

// TellerName 名言の発案者名。誰が言ったか
type TellerName string

// Sentence 名言の内容。何を言ったか
type Sentence string
