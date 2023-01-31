package infrastructure

import types "github.com/isteca85/car-pooling-challenge/pkg/domain"

type DataBase interface {
	OpenDdbb() error
	CloseDdbb()
	CleanDdbb() error
	InsertCar(car types.Car) (int64, error)
	InsertJourney(journey types.Journey) (int64, error)
	GetFreeCar(people int) (int64, error)
	GetNextJourneyByCar(car types.Car) (types.JourneyDB, error)
	UpdateCarByJourney(idCar int64, idJourney int64) error
	UnlinkCarsByJourney(idJourney int64) error
	RemoveJourney(journey types.JourneyDB) error
	GetCar(id int64) (types.Car, error)
	GetCarByJourney(id int64) (types.Car, error)
	GetJourney(id int64) (types.JourneyDB, error)
}
