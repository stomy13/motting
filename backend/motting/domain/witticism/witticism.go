package witticism

import (
	"errors"

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

func NewTellerName(newTellerName string) (*TellerName, error) {
	tellerName := TellerName(newTellerName)
	return &tellerName, tellerName.valid()
}

// 1文字以上であることの確認
func (newTellerName *TellerName) valid() error {
	if len(string(*newTellerName)) <= 0 {
		return errors.New("tellername must not be empty.")
	}
	return nil
}

// Sentence 名言の内容。何を言ったか
type Sentence string

func NewSentence(newSentence string) (*Sentence, error) {
	sentence := Sentence(newSentence)
	return &sentence, sentence.valid()
}

// 1文字以上であることの確認
func (newSentence *Sentence) valid() error {
	if len(string(*newSentence)) <= 0 {
		return errors.New("sentence must not be empty.")
	}
	return nil
}
