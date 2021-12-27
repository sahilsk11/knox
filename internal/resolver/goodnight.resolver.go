package resolver

import (
	"encoding/json"
)

func (m httpServer) goodnight([]byte) ([]byte, error) {
	err := m.LightsApp.GoodnightScene()
	if err != nil {
		return nil, err
	}

	response := map[string]string{
		"success": "true",
	}
	responseBody, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
