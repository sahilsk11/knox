package thermostat

type ThermostatName string

const (
	Thermostat_SahilBedroom ThermostatName = "SAHIL_THERMOSTAT"
	Thermostat_GuestBedroom ThermostatName = "GUEST_THERMOSTAT"
	Thermostat_LivingRoom   ThermostatName = "LIVING_THERMOSTAT"
)

type Thermostat struct {
	ThermostatName          ThermostatName
	HomeAssistantEntityName string `json:"homeAssistantEntityName"`
}
