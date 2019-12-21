package v1

import (
	"context"
	"github.com/djamboe/mtools-location-service/controllers"
	"github.com/djamboe/mtools-location-service/infrastructures"
	"github.com/djamboe/mtools-location-service/repositories"
	"github.com/djamboe/mtools-location-service/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
)

type IserviceContainer interface {
	InjectLocationController() controllers.LocationController
}

type kernel struct{}

func (k *kernel) InjectLocationController() controllers.LocationController {
	//mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/marketing-tools?charset=utf8")
	//mysqlHandler := &infrastructures.DBHandler{}
	//mysqlHandler.Conn = mysqlConn
	//loginRepository := &repositories.LoginRepository{mysqlHandler}
	//loginService := &services.LoginService{&repositories.LoginRepositoryWithCircuitBreaker{loginRepository}}
	//loginController := controllers.LoginController{loginService}

	//test mongodb
	c := GetClient()
	mongoCon := c.Ping(context.Background(), readpref.Primary())
	if mongoCon != nil {
		log.Fatal("Couldn't connect to the database", mongoCon)
	}

	mongoDBConn := c
	mongoDBHandler := &infrastructures.MongoDBHandler{}
	mongoDBHandler.Conn = mongoDBConn
	locationRepository := &repositories.LocationRepository{mongoDBHandler}
	locationService := &services.LocationService{&repositories.LocationRepositoryWithCircuitBreaker{locationRepository}}
	locationController := controllers.LocationController{locationService}
	//test mongodb

	return locationController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IserviceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
