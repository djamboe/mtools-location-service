package controllers

import (
	"github.com/djamboe/mtools-location-service/interfaces"
	"github.com/djamboe/mtools-location-service/models"
)

type LocationController struct {
	interfaces.ILocationRepository
}

func (controller *LocationController) RecordLocation(param models.LocationParamModel) error {
	err := controller.InsertNewLocation(param)
	if err != nil {
		panic(err)
	}
	return nil
}
