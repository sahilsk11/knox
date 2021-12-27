package light_controller

type LightName string

// decided to not list all the enum types and validate in the db instead

// const (
// 	Light_Sahil_BedLight LightName = "SAHIL_BEDLight"
// 	Light_Living         LightName = "LIVING"
// 	Light_Kitchen        LightName = "KITCHEN"
// 	Light_Dining         LightName = "DINING_LIGHT"
// 	Light_Porch          LightName = "PORCH"
// 	Light_Entry          LightName = "ENTRY"
// 	Light_Guest_Closet   LightName = "GUEST_CLOSET"
// 	Light_Guest_Bathroom LightName = "GUEST_BATHROOM"
// 	Light_Guest_Bedroom  LightName = "GUEST_BEDROOM"
// 	Light_Sahil_Closet   LightName = "SAHIL_CLOSET"

// )

type LightSwitchType string

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

type Light struct {
	LightName               LightName
	SwitchType              LightSwitchType `json:"switchType"`
	HomeAssistantEntityName string          `json:"homeAssistantEntityName"`
}
