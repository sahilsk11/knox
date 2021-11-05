package service

type ApplianceService interface{}

type applianceService struct{}

func NewApplianceService() ApplianceService {
	return applianceService{}
}
