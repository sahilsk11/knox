package service

import (
	"fmt"

	"github.com/pkg/errors"
	domain "github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/util"
)

type PlayerService interface {
	ListAvailableDevices() ([]domain.PlayerDevice, error)
	ListGenres() ([]domain.PlaybackGenre, error)
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

func (m playerService) ListGenres() ([]domain.PlaybackGenre, error) {
	return domain.ListGenreTypes(), nil
}

type StartPlaybackInput struct {
	DeviceFilter DeviceFilter
	Genre        domain.PlaybackGenre
}

type DeviceFilter struct {
	DeviceID            *string
	DeviceNameSimilarTo *string
}

func (m playerService) resolveDeviceFilter(filter DeviceFilter) (*string, error) {
	if filter.DeviceID != nil {
		return filter.DeviceID, nil
	} else if filter.DeviceNameSimilarTo != nil {
		return nil, fmt.Errorf("invalid device filter")
	}

	devices, err := m.PlayerVendorRepository.ListDevices()
	if err != nil {
		return nil, fmt.Errorf("[service] failed to list devices: %s", err.Error())
	} else if len(devices) == 0 {
		return nil, fmt.Errorf("[service] no devices returned")
	}

	// pick the closest device
	deviceID := findClosestDevice(*filter.DeviceNameSimilarTo, devices)

	return &deviceID, nil
}

func (m playerService) StartPlayback(input StartPlaybackInput) error {
	deviceIDPtr, err := m.resolveDeviceFilter(input.DeviceFilter)
	if err != nil {
		return err
	}
	deviceID := *deviceIDPtr

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
