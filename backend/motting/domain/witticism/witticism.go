package witticism

import (
	"errors"

	"github.com/MasatoTokuse/motting/motting/domain/user"
	"github.com/MasatoTokuse/motting/motting/error_response"
	"github.com/google/uuid"
)

// Witticism 名言。より良い人生をつくるための言葉
type Witticism struct {
	Id         *WitticismId
	TellerName *TellerName
	Sentence   *Sentence
	OwnerId    *user.UserId
}

func NewWitticism(
	tellerNameStr string,
	sentenceStr string,
	ownerId *user.UserId,
) (*Witticism, error) {
	validateErrors := error_response.NewValidateErrors()

	witticismId, err := NewWitticismId()
	validateErrors.Append("witticismId", err)

	tellerName, err := NewTellerName(tellerNameStr)
	validateErrors.Append("tellerName", err)

	sentence, err := NewSentence(sentenceStr)
	validateErrors.Append("sentence", err)

	if validateErrors.HasError() {
		return nil, validateErrors
	}

	return &Witticism{
		Id:         witticismId,
		TellerName: tellerName,
		Sentence:   sentence,
		OwnerId:    ownerId,
	}, nil
}

func NewWitticismWithUUID(
	id string,
	tellerNameStr string,
	sentenceStr string,
	ownerIdStr string,
) (*Witticism, error) {
	validateErrors := error_response.NewValidateErrors()

	witticismId, err := NewWitticismIdFromString(id)
	validateErrors.Append("witticismId", err)

	tellerName, err := NewTellerName(tellerNameStr)
	validateErrors.Append("tellerName", err)

	sentence, err := NewSentence(sentenceStr)
	validateErrors.Append("sentence", err)

	ownerId := user.UserId(ownerIdStr)

	if validateErrors.HasError() {
		return nil, validateErrors
	}

	return &Witticism{
		Id:         witticismId,
		TellerName: tellerName,
		Sentence:   sentence,
		OwnerId:    &ownerId,
	}, nil
}

// WitticismId 名言のID。このドメインのID
type WitticismId string

func NewWitticismId() (*WitticismId, error) {
	witticism := WitticismId(uuid.NewString())
	return &witticism, witticism.valid()
}

func NewWitticismIdFromString(id string) (*WitticismId, error) {
	witticism := WitticismId(id)
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
