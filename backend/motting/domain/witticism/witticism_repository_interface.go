package witticism

type WitticismRepositoryInterface interface {
	// TODO:後で実装する
	// AllWitticism(userId *user.UserId) error
	Save(witticism *Witticism) error
	// Delete(witticismId *WitticismId) error
}
