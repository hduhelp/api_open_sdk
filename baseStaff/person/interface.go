package person

type InfoInterface interface {
	GetPersonInfo() (*Info, error)
}
