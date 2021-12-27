package service

import (
	domain "github.com/sahilsk11/knox/internal/domain/thermostat"
	"github.com/sahilsk11/knox/internal/repository"
)

type ThermostatService interface {
	TurnOff(thermostatName domain.ThermostatName) error
	SetTemperature(SetTemperatureInput) error
}

func NewThermostatService(thermostatRepository domain.ThermostatRepository, thermostatDatabaseRepository repository.ThermostatDatabaseRepository) ThermostatService {
	return thermostatService{
		ThermostatRepository:         thermostatRepository,
		ThermostatDatabaseRepository: thermostatDatabaseRepository,
	}
}

type thermostatService struct {
	ThermostatRepository         domain.ThermostatRepository
	ThermostatDatabaseRepository repository.ThermostatDatabaseRepository
}

func (m thermostatService) TurnOff(thermostatName domain.ThermostatName) error {
	thermostat, err := m.ThermostatDatabaseRepository.GetThermostat(thermostatName)
	if err != nil {
		return err
	}

	return m.ThermostatRepository.TurnOffThermostat(thermostat.HomeAssistantEntityName)
}

type SetTemperatureInput struct {
	ThermostatName    domain.ThermostatName
	TargetTemperature int
}

func (m thermostatService) SetTemperature(input SetTemperatureInput) error {
	thermostat, err := m.ThermostatDatabaseRepository.GetThermostat(input.ThermostatName)
	if err != nil {
		return err
	}

	setThermostatInput := domain.SetTemperatureInput{
		Thermostat:        *thermostat,
		TargetTemperature: input.TargetTemperature,
	}
	return m.ThermostatRepository.SetTemperature(setThermostatInput)
}
