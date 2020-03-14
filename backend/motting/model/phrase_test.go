package model

import (
	"testing"
)

func TestPhrase(t *testing.T) {
	phrase := &Phrase{
		UserID: "",
		Text:   "",
		Author: "",
	}
	t.Log(phrase.UserID)
	t.Log(phrase.Text)
	t.Log(phrase.Author)
}
