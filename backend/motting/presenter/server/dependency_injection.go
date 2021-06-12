package server

import (
	usercase_witticism "github.com/MasatoTokuse/motting/motting/application/witticism"
	domain_witticism "github.com/MasatoTokuse/motting/motting/domain/witticism"
	persistance "github.com/MasatoTokuse/motting/motting/infrastracture/persistence/mysql"
	presenter_witticism "github.com/MasatoTokuse/motting/motting/presenter/witticism"
	"gorm.io/gorm"
)

type Compornents struct {
	WitticismRepository  domain_witticism.WitticismRepositoryInterface
	WitticismUsecase     usercase_witticism.WitticismUsecaseInterface
	WitticismContoroller *presenter_witticism.WitticismController
}

func NewCompornents(connection *gorm.DB) Compornents {
	var repository = persistance.NewWitticismRepository(connection)
	var usecase = usercase_witticism.NewWitticismUsecaseInteractor(repository)
	var controller = presenter_witticism.NewWitticismController(usecase)
	var compornents = Compornents{WitticismRepository: repository, WitticismUsecase: usecase, WitticismContoroller: controller}
	return compornents
}
