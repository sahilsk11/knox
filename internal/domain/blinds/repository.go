package blinds

type BlindsControllerRepository interface {
	SetHeight(BlindsName, int) error
	GetHeight(BlindsName) (int, error)
}
