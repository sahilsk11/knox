package app

type BlindsApp interface {
	CloseAll() error
	OpenAll() error
	OpenEntry() error
	CloseEntry() error
}
