package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/djamboe/mtools-location-service/interfaces"
	"github.com/djamboe/mtools-location-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type LocationRepositoryWithCircuitBreaker struct {
	LocationRepository interfaces.ILocationRepository
}

type LocationRepository struct {
	//interfaces.IDbHandler
	interfaces.IMongoDBHandler
}

func (repository *LocationRepositoryWithCircuitBreaker) InsertNewLocation(param models.LocationParamModel) error {
	output := make(chan error, 1)
	hystrix.ConfigureCommand("create_post", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("create_post", func() error {
		locationData := repository.LocationRepository.InsertNewLocation(param)
		output <- locationData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out
	case err := <-errors:
		println(err)
		return err
	}
}

func (repository *LocationRepository) InsertNewLocation(param models.LocationParamModel) error {
	var p models.Point
	p.UserId = param.UserId
	p.Location = models.NewPoint(param.Lat, param.Long)
	p.MemberName = param.MemberName
	p.Description = param.Description
	err := repository.InsertOne(p, "location_log", "maroon_martools")
	if err != nil {
		panic(err)
	}
	return nil
}
