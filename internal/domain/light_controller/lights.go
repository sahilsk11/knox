package light_controller

type LightName string

const (
	Light_Sahil_Bedroom       LightName = "SAHIL_BEDROOM"
	Light_Living              LightName = "LIVING"
	Light_Kitchen             LightName = "KITCHEN"
	Light_Dining              LightName = "DINING"
	Light_Porch               LightName = "PORCH"
	Light_Entry               LightName = "ENTRY"
	Light_Guest_Closet        LightName = "GUEST_CLOSET"
	Light_Guest_Bathroom      LightName = "GUEST_BATHROOM"
	Light_Guest_Bedroom       LightName = "GUEST_BEDROOM"
	Light_Sahil_Closet        LightName = "SAHIL_CLOSET"
	Light_Sahil_Bathroom      LightName = "SAHIL_BATHROOM"
	Light_Downstairs_Bathroom LightName = "DOWNSTAIRS_BATHROOM"
	Light_Staircase           LightName = "STAIRCASE"
	Light_Storage             LightName = "STORAGE"
	Light_Kitchen_Accent      LightName = "KITCHEN_ACCENT"
)

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
