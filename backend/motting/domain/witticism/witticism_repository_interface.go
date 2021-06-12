package witticism

import "github.com/MasatoTokuse/motting/motting/domain/user"

type WitticismRepositoryInterface interface {
	FindByOwnerId(ownerId *user.UserId) ([]*Witticism, error)
	Save(witticism *Witticism) error
	// Delete(witticismId *WitticismId) error
}
