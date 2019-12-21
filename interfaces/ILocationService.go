package interfaces

import "github.com/djamboe/mtools-location-service/models"

type ILocationService interface {
	CreateLocation(param models.LocationParamModel) error
}
