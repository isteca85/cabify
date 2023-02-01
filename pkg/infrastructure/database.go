package infrastructure

import types "github.com/isteca85/car-pooling-challenge/pkg/domain"

type DataBase interface {
	OpenDdbb() error
	CloseDdbb()
	CleanDdbb() error
	InsertCar(car types.Car) (int64, error)
	InsertJourney(journey types.Journey) (int64, error)
	GetFreeCar(people int) (int64, error)
	GetNextJourneyBySeats(seats int) (types.JourneyDB, error)
	UpdateCarByJourney(idCar int64, idJourney int64) error
	UnlinkCarsByJourney(idJourney int64) (int64, int, error)
	RemoveJourney(journey types.JourneyDB) error
	GetCar(id int64) (types.Car, error)
	GetJourney(id int64) (types.JourneyDB, error)
	UpdateJourneyStatus(journey int64) error
	GetCarByJourney(id int64) (types.Car, error)
}
