package model

import (
	"testing"
)

func TestSubscription(t *testing.T) {
	sub := &Subscription{
		UserID:   "whitebox",
		Endpoint: "ep",
		P256dh:   "P256dhPOUYG",
		Auth:     "AUTHEIOIJTES",
	}
	t.Log(sub.UserID)
	t.Log(sub.Endpoint)
	t.Log(sub.P256dh)
	t.Log(sub.Auth)
}
