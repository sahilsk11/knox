package player

type Player struct{}

type PlayerDevice struct {
	DeviceName string `json:"device_name"`
	DeviceID   string `json:"device_id"`
}
