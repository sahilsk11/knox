package service

type PlayerService interface{}

type playerService struct{}

func NewPlayerService() PlayerService {
	return playerService{}
}
