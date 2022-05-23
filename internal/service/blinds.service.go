package service

import "github.com/sahilsk11/knox/internal/domain/blinds"

type BlindsService interface {
	OpenOne(blinds.BlindsName) error
	OpenMany([]blinds.BlindsName) error
	CloseMany([]blinds.BlindsName) error
	CloseOne(blinds.BlindsName) error
}
