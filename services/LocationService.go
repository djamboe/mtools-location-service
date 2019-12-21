package services

import (
	"github.com/djamboe/mtools-location-service/interfaces"
	"github.com/djamboe/mtools-location-service/models"
)

type LocationService struct {
	interfaces.ILocationRepository
}

func (service *LocationService) CreateLocation(param models.LocationParamModel) error {
	err := service.InsertNewLocation(param)
	if err != nil {
		panic(err)
	}
	return nil
}
