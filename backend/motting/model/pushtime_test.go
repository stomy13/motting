package model

import (
	"testing"
)

func TestPushTime(t *testing.T) {
	pt := &PushTime{
		UserID: "User1",
		PushAt: "10:10",
	}
	t.Log(pt.UserID)
	t.Log(pt.PushAt)
}
