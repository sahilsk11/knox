package blinds

type BlindsControllerRepository interface {
	SetHeight(map[BlindsName]int) error
	GetHeight(BlindsName) (int, error)
}
