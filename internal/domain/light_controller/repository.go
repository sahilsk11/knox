package light_controller

type LightControllerRepository interface {
	ControlLights(input ControlLightsInput) error
}

type ControlLightsInput struct {
	Light      Light
	Brightness *int
	State      LightState
}
