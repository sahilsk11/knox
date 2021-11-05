package player

type PlayerVendorRepository interface {
	ListDevices() ([]PlayerDevice, error)
	StartPlayback(StartPlaybackInput) error
}

type StartPlaybackInput struct {
	DeviceID string
}
