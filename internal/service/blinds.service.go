package service

import domain "github.com/sahilsk11/knox/internal/domain/blinds"

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
