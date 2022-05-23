package service

import (
	"fmt"

	domain "github.com/sahilsk11/knox/internal/domain/blinds"
)

type BlindsService interface {
	OpenOne(domain.BlindsName) error
	OpenMany([]domain.BlindsName) error
	CloseMany([]domain.BlindsName) error
	CloseOne(domain.BlindsName) error
	ListBlinds() []domain.BlindsName
}

type blindsService struct {
	BlindsRepository domain.BlindsControllerRepository
}

func (m blindsService) ListBlinds() []domain.BlindsName {
	return []domain.BlindsName{
		domain.Blinds_BigWindowFirstThird,
		domain.Blinds_BigWindowSecondThirds,
		domain.Blinds_Door,
		domain.Blinds_LivingRoom,
	}
}

func (m blindsService) OpenOne(name domain.BlindsName) error {
	err := m.BlindsRepository.SetHeight(name, 0)
	if err != nil {
		return fmt.Errorf("open_one_blind_fail: %w", err)
	}
	return nil
}

func (m blindsService) CloseOne(name domain.BlindsName) error {
	err := m.BlindsRepository.SetHeight(name, 100)
	if err != nil {
		return fmt.Errorf("close_one_blind_fail: %w", err)
	}
	return nil
}
