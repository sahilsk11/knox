package service

import (
	"fmt"

	"github.com/pkg/errors"
	domain "github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/util"
)

type PlayerService interface {
	ListAvailableDevices() ([]domain.PlayerDevice, error)
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

func (m playerService) ListAvailableDevices() ([]domain.PlayerDevice, error) {
	devices, err := m.PlayerVendorRepository.ListDevices()
	if err != nil {
		return nil, errors.Wrap(err, "failed to list devices")
	}

	return devices, nil
}

type StartPlaybackInput struct {
	DeviceNameSimilarTo string
	Genre               domain.PlaybackGenre
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
		Genre:    input.Genre,
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
