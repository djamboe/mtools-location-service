package interfaces

import (
	"github.com/djamboe/mtools-location-service/models"
)

type ILocationRepository interface {
	InsertNewLocation(model models.LocationParamModel) error
}
