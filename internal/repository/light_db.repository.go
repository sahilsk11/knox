package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	domain "github.com/sahilsk11/knox/internal/domain/light_controller"
)

type LightDatabaseRepository interface {
	GetRoom(domain.LightName) (*domain.Light, error)
}

type lightDatabaseRepository struct {
	LightDatabase map[domain.LightName]domain.Light
}

type lightDatabase struct {
	Comments []string                          `json:"comments"`
	Lights   map[domain.LightName]domain.Light `json:"lights"`
}

func NewLightDatabaseRepository(filepath string) (LightDatabaseRepository, error) {
	m := lightDatabase{}

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

	return lightDatabaseRepository{
		LightDatabase: m.Lights,
	}, nil
}

func (m lightDatabaseRepository) GetRoom(LightName domain.LightName) (*domain.Light, error) {
	if room, ok := m.LightDatabase[LightName]; ok {
		room.LightName = LightName
		return &room, nil
	}

	return nil, fmt.Errorf("key %s does not exist in light db", string(LightName))
}
