package witticism

import (
	"github.com/MasatoTokuse/motting/motting/domain/user"
)

type WitticismRepositoryInterface interface {
	allWitticism(userId *user.UserId)
	save(witticism *Witticism)
	delete(witticismId *WitticismId)
}
