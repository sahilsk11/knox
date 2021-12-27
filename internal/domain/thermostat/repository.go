package thermostat

type ThermostatRepository interface {
	SetTemperature(SetTemperatureInput) error
	GetTemperature(entityName string) error
	TurnOff(entityName string) error
}

type SetTemperatureInput struct {
	Thermostat        Thermostat
	TargetTemperature int
}
