package app

import (
	"fmt"

	domain "github.com/sahilsk11/knox/internal/domain/blinds"
	"github.com/sahilsk11/knox/internal/service"
)

type BlindsApp interface {
	CloseAll() error
	OpenAll() error
	OpenEntry() error
	CloseEntry() error
}

type blindsApp struct {
	BlindsService service.BlindsService
}

func (m blindsApp) CloseAll() error {
	blinds := m.BlindsService.ListBlinds()
	err := m.BlindsService.CloseMany(blinds)
	if err != nil {
		return fmt.Errorf("close_all_blinds_fail: %w", err)
	}
	return nil
}

func (m blindsApp) OpenAll() error {
	blinds := m.BlindsService.ListBlinds()
	err := m.BlindsService.OpenMany(blinds)
	if err != nil {
		return fmt.Errorf("open_all_blinds_fail: %w", err)
	}
	return nil
}

func (m blindsApp) OpenEntry() error {
	err := m.BlindsService.OpenOne(domain.Blinds_Door)
	if err != nil {
		return fmt.Errorf("open_entry_blinds_fail: %w", err)
	}
	return nil
}

func (m blindsApp) CloseEntry() error {
	err := m.BlindsService.CloseOne(domain.Blinds_Door)
	if err != nil {
		return fmt.Errorf("close_entry_blinds_fail: %w", err)
	}
	return nil
}
