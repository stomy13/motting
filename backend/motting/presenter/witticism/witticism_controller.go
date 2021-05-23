package witticism

import (
	"net/http"

	"github.com/MasatoTokuse/motting/motting/application/witticism"
	"github.com/MasatoTokuse/motting/motting/presenter/internal"
)

type WitticismController struct {
	WitticismUsecase witticism.WitticismUsecaseInterface
}

func NewWitticismController(witticismUsecase witticism.WitticismUsecaseInterface) *WitticismController {
	return &WitticismController{
		WitticismUsecase: witticismUsecase,
	}
}

func (controller *WitticismController) AddWitticism(response http.ResponseWriter, request *http.Request) {
	var err error
	defer func() {
		internal.Dispatch(err, response)
	}()

	var addWitticismCommand witticism.AddWitticismCommand
	err = internal.UnmarshalJson(request.Body, &addWitticismCommand)
	if err != nil {
		return
	}
	err = controller.WitticismUsecase.AddWitticism(&addWitticismCommand)
}
