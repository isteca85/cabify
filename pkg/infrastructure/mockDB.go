package infrastructure

import (
	types "github.com/isteca85/car-pooling-challenge/pkg/domain"
)

type mockDB struct {
	journeysDB []types.JourneyDB
	carsDB     []types.CarDB
}

func (m *mockDB) OpenDdbb() error {
	return nil
}

func (m *mockDB) CloseDdbb() {

}

func (m *mockDB) CleanDdbb() error {
	m.carsDB = nil
	m.journeysDB = nil
	return nil
}

func (m *mockDB) InsertCar(car types.Car) (int64, error) {
	var carDB = types.CarDB{ID: car.ID, Seats: car.Seats, ID_Journey: 0}
	m.carsDB = append(m.carsDB, carDB)
	return 0, nil
}

func (m *mockDB) InsertJourney(journey types.Journey) (int64, error) {
	m.journeysDB = append(m.journeysDB, types.JourneyDB{ID: journey.ID, People: journey.People, Status: types.Unasigned})
	return 0, nil
}

func (m *mockDB) GetFreeCar(people int) (int64, error) {
	for _, car := range m.carsDB {
		if car.Seats >= people && car.ID_Journey == 0 {
			return car.ID, nil
		}
	}
	return 0, nil
}

func (m *mockDB) GetNextJourneyByCar(car types.Car) (types.JourneyDB, error) {
	for _, journey := range m.journeysDB {
		if journey.Status == types.Unasigned && journey.People <= car.Seats {
			return journey, nil
		}
	}
	return types.JourneyDB{}, nil
}

func (m *mockDB) UpdateCarByJourney(idCar int64, idJourney int64) error {
	for i, car := range m.carsDB {
		if car.ID == idCar {
			m.carsDB[i].ID_Journey = idJourney
			return nil
		}
	}
	return nil
}

func (m *mockDB) UnlinkCarsByJourney(idJourney int64) error {
	for i, car := range m.carsDB {
		if car.ID_Journey == idJourney {
			m.carsDB[i].ID_Journey = 0
		}
	}
	return nil
}

func (m *mockDB) RemoveJourney(journey types.JourneyDB) error {
	for i, journeyDB := range m.journeysDB {
		if journeyDB.ID == journey.ID {
			m.journeysDB = append(m.journeysDB[:i], m.journeysDB[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *mockDB) GetCar(id int64) (types.Car, error) {
	for _, car := range m.carsDB {
		if car.ID == id {
			return types.Car{ID: car.ID, Seats: car.Seats}, nil
		}
	}
	return types.Car{}, nil
}

func (m *mockDB) GetCarByJourney(id int64) (types.Car, error) {
	for _, car := range m.carsDB {
		if car.ID_Journey == id {
			return types.Car{ID: car.ID, Seats: car.Seats}, nil
		}
	}
	return types.Car{}, nil
}

func (m *mockDB) GetJourney(id int64) (types.JourneyDB, error) {
	for _, journey := range m.journeysDB {
		if journey.ID == id {
			return journey, nil
		}
	}
	return types.JourneyDB{}, nil
}
