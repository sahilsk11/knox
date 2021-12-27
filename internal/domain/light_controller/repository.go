package light_controller

type LightControllerRepository interface {
	ControlLights(input ControlLightsInput) error
}

type ControlLightsInput struct {
	EntityName string
	Brightness *int
	State      LightState
}
