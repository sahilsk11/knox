package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	domain "github.com/sahilsk11/knox/internal/domain/thermostat"
)

type ThermostatDatabaseRepository interface {
	GetThermostat(domain.ThermostatName) (*domain.Thermostat, error)
}

type thermostatDatabaseRepository struct {
	ThermostatDatabase map[domain.ThermostatName]domain.Thermostat
}

type thermostatDatabase struct {
	Comments    []string                                    `json:"comments"`
	Thermostats map[domain.ThermostatName]domain.Thermostat `json:"thermostats"`
}

func NewThermostatDatabaseRepository(filepath string) (ThermostatDatabaseRepository, error) {
	m := thermostatDatabase{}

	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open light db file: %s", err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read light db file: %s", err.Error())
	}

	err = json.Unmarshal(byteValue, &m)
	if err != nil {
		return nil, fmt.Errorf("failed to load light db: %s", err.Error())
	}

	return thermostatDatabaseRepository{
		ThermostatDatabase: m.Thermostats,
	}, nil
}

func (m thermostatDatabaseRepository) GetThermostat(ThermostatName domain.ThermostatName) (*domain.Thermostat, error) {
	if room, ok := m.ThermostatDatabase[ThermostatName]; ok {
		room.ThermostatName = ThermostatName
		return &room, nil
	}

	return nil, fmt.Errorf("key %s does not exist in thermostat db", string(ThermostatName))
}
