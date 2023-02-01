package infrastructure

import (
	"errors"
	types "github.com/isteca85/car-pooling-challenge/pkg/domain"
)

type MockDB struct {
	JourneysDB []types.JourneyDB
	CarsDB     []types.CarDB
}

func (m *MockDB) OpenDdbb() error {
	return nil
}

func (m *MockDB) CloseDdbb() {

}

func (m *MockDB) CleanDdbb() error {
	m.CarsDB = nil
	m.JourneysDB = nil
	return nil
}

func (m *MockDB) InsertCar(car types.Car) (int64, error) {
	var carDB = types.CarDB{ID: car.ID, Seats: car.Seats, ID_Journey: 0}
	m.CarsDB = append(m.CarsDB, carDB)
	return 0, nil
}

func (m *MockDB) InsertJourney(journey types.Journey) (int64, error) {
	m.JourneysDB = append(m.JourneysDB, types.JourneyDB{ID: journey.ID, People: journey.People, Status: types.Unasigned})
	return 0, nil
}

func (m *MockDB) GetFreeCar(people int) (int64, error) {
	for _, car := range m.CarsDB {
		if car.Seats >= people && car.ID_Journey == 0 {
			return car.ID, nil
		}
	}
	return 0, nil
}

func (m *MockDB) GetNextJourneyBySeats(seats int) (types.JourneyDB, error) {
	for _, journey := range m.JourneysDB {
		if journey.Status == types.Unasigned && journey.People <= seats {
			return journey, nil
		}
	}
	return types.JourneyDB{}, nil
}

func (m *MockDB) UpdateCarByJourney(idCar int64, idJourney int64) error {
	for i, car := range m.CarsDB {
		if car.ID == idCar {
			m.CarsDB[i].ID_Journey = idJourney
			return nil
		}
	}
	return nil
}

func (m *MockDB) UnlinkCarsByJourney(idJourney int64) (int64, int, error) {
	var carID int64
	var seats int
	for i, car := range m.CarsDB {
		if car.ID_Journey == idJourney {
			m.CarsDB[i].ID_Journey = 0
			carID = car.ID
			seats = car.Seats
		}
	}
	return carID, seats, nil
}

func (m *MockDB) RemoveJourney(journey types.JourneyDB) error {
	for i, journeyDB := range m.JourneysDB {
		if journeyDB.ID == journey.ID {
			m.JourneysDB = append(m.JourneysDB[:i], m.JourneysDB[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockDB) GetCar(id int64) (types.Car, error) {
	for _, car := range m.CarsDB {
		if car.ID == id {
			return types.Car{ID: car.ID, Seats: car.Seats}, nil
		}
	}
	return types.Car{}, nil
}

func (m *MockDB) GetJourney(id int64) (types.JourneyDB, error) {
	for _, journey := range m.JourneysDB {
		if journey.ID == id {
			return journey, nil
		}
	}
	return types.JourneyDB{}, errors.New("Journey not found")
}

func (m *MockDB) UpdateJourneyStatus(journey int64) error {
	for i, journeyDB := range m.JourneysDB {
		if journeyDB.ID == journey {
			m.JourneysDB[i].Status = types.Asigned
			return nil
		}
	}
	return nil
}

func (m *MockDB) GetCarByJourney(id int64) (types.Car, error) {
	for _, car := range m.CarsDB {
		if car.ID_Journey == id {
			return types.Car{ID: car.ID, Seats: car.Seats}, nil
		}
	}
	return types.Car{}, nil
}
