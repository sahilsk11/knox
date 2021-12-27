package thermostat

type ThermostatRepository interface {
	SetTemperature(SetTemperatureInput) error
	GetTemperature(entityName string) (*int, error)
	TurnOffThermostat(entityName string) error
}

type SetTemperatureInput struct {
	Thermostat        Thermostat
	TargetTemperature int
}
