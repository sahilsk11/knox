package light_controller

type RoomName string

const (
	Room_Sahil_Bedroom RoomName = "SAHIL_BEDROOM"
	Room_Living        RoomName = "LIVING"
	Room_Kitchen       RoomName = "KITCHEN"
)

type RoomSwitchType string

const (
	// a switch is a boolean on/off controller
	SwitchType_Switch = "SWITCH"
	// adjustable switches may have a range of valid brightness
	SwitchType_Adjustable = "ADJUSTABLE"
)

type LightState string

const (
	LightState_On  LightState = "ON"
	LightState_Off LightState = "OFF"
)
