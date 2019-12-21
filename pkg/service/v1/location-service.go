package v1

import (
	"context"
	"github.com/djamboe/mtools-location-service/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/djamboe/mtools-location-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type locationServiceServer struct {
}

// NewToDoServiceServer creates ToDo service
func NewLocationServiceServer() v1.LocationServiceServer {
	return &locationServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *locationServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	} else {
		return status.Errorf(codes.Unimplemented,
			"please input your api version")
	}
	return nil
}

// Create new todo task
func (s *locationServiceServer) CreateLocation(ctx context.Context, req *v1.LocationRequest) (*v1.LocationResponse, error) {
	// check if the API version requested by client is supported by server
	var locationParam models.LocationParamModel
	message := "Successfully Create Location !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	locationParam.UserId = req.UserId
	locationParam.Lat = req.Lat
	locationParam.Long = req.Long
	locationParam.Description = req.Description
	locationParam.MemberName = req.MemberName

	locationController := ServiceContainer().InjectLocationController()
	err := locationController.InsertNewLocation(locationParam)

	if err != nil {
		message = "Failed insert location, an error occured"
		errorStatus = true
	}

	return &v1.LocationResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
	}, nil
}
