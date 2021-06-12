package witticism

type WitticismUsecaseInterface interface {
	AddWitticism(command *AddWitticismCommand) error
}

type AddWitticismCommand struct {
	TellerName string `json:"tellerName"`
	Sentence   string `json:"sentence"`
	OwnerId    string `json:"ownerId"`
}
