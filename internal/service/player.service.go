package service

import (
	"fmt"

	domain "github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/util"
)

type PlayerService interface {
	StartPlayback(input StartPlaybackInput) error
}

type playerService struct {
	PlayerVendorRepository domain.PlayerVendorRepository
}

func NewPlayerService(playerVendorRepository domain.PlayerVendorRepository) PlayerService {
	return playerService{
		PlayerVendorRepository: playerVendorRepository,
	}
}

type StartPlaybackInput struct {
	DeviceNameSimilarTo string
}

func (m playerService) StartPlayback(input StartPlaybackInput) error {
	devices, err := m.PlayerVendorRepository.ListDevices()
	if err != nil {
		return fmt.Errorf("[service] failed to list devices: %s", err.Error())
	} else if len(devices) == 0 {
		return fmt.Errorf("[service] no devices returned")
	}

	// pick the closest device
	deviceID := findClosestDevice(input.DeviceNameSimilarTo, devices)

	playbackInput := domain.StartPlaybackInput{
		DeviceID: deviceID,
	}
	err = m.PlayerVendorRepository.StartPlayback(playbackInput)
	if err != nil {
		return fmt.Errorf("[service] failed to start playback: %s", err.Error())
	}

	return nil
}

func findClosestDevice(pattern string, devices []domain.PlayerDevice) string {
	maxScore := float32(0)
	deviceID := ""
	for _, device := range devices {
		score := util.CompareTwoStrings(pattern, device.DeviceName)
		if score > maxScore {
			maxScore = score
			deviceID = device.DeviceID
		}
	}
	return deviceID
}
