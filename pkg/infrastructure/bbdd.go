package infrastructure

import (
	"database/sql"
	"fmt"
	types "github.com/isteca85/car-pooling-challenge/pkg/domain"
)

type Bbdd struct {
	db *sql.DB
}

func (b *Bbdd) OpenDdbb() error {
	usuario := "root"
	pass := "RoadToSanSil.2309."
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "cabify"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return err
	}
	b.db = db
	return nil
}

func (b *Bbdd) CloseDdbb() {
	b.db.Close()
}

func (b *Bbdd) CleanDdbb() error {
	_, err := b.db.Exec("TRUNCATE TABLE cabify.cars")
	if err != nil {
		return err
	}
	_, err = b.db.Exec("DELETE FROM cabify.journeys")
	if err != nil {
		return err
	}
	return nil
}

func (b *Bbdd) InsertCar(car types.Car) (int64, error) {
	stmt, err := b.db.Prepare("INSERT INTO cabify.cars (ID, seats) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(car.ID, car.Seats)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (b *Bbdd) InsertJourney(journey types.Journey) (int64, error) {
	stmt, err := b.db.Prepare("INSERT INTO cabify.journeys (ID, people, status) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(journey.ID, journey.People, types.Unasigned)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (b *Bbdd) GetFreeCar(people int) (int64, error) {
	var idCar int64
	err := b.db.QueryRow("SELECT ID FROM cabify.cars WHERE ID_Journey IS NULL AND seats >= ? ORDER BY seats LIMIT 1", people).Scan(&idCar)
	if err != nil {
		return idCar, err
	}
	return idCar, nil
}

func (b *Bbdd) GetNextJourneyBySeats(seats int) (types.JourneyDB, error) {
	var journey types.JourneyDB
	err := b.db.QueryRow("SELECT ID, people, status FROM cabify.journeys WHERE people <= ? AND status = ? ORDER BY time", seats, types.Unasigned).Scan(&journey.ID, &journey.People, &journey.Status)
	if err != nil {
		return journey, err
	}
	return journey, nil
}

func (b *Bbdd) UpdateCarByJourney(idCar int64, idJourney int64) error {
	err := b.db.QueryRow("UPDATE cabify.journeys SET status = ? WHERE ID = ? ", types.Asigned, idJourney).Err()
	if err != nil {
		return err
	}
	err = b.db.QueryRow("UPDATE cabify.cars SET ID_Journey = ? WHERE ID = ?", idJourney, idCar).Err()
	if err != nil {
		return err
	}
	return nil
}

func (b *Bbdd) UnlinkCarsByJourney(idJourney int64) (int64, int, error) {
	var idCar int64
	var seats int
	err := b.db.QueryRow("SELECT ID, seats FROM cabify.cars WHERE ID_Journey = ?", idJourney).Scan(&idCar, &seats)
	if err != nil {
		return 0, 0, err
	}
	err = b.db.QueryRow("UPDATE cabify.cars SET ID_Journey = NULL WHERE ID_Journey = ? ", idJourney).Err()
	if err != nil {
		return 0, 0, err
	}
	return idCar, seats, nil
}

func (b *Bbdd) RemoveJourney(journey types.JourneyDB) error {
	err := b.db.QueryRow("DELETE FROM cabify.journeys WHERE ID = ? ", journey.ID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (b *Bbdd) GetCar(id int64) (types.Car, error) {
	var car types.Car
	err := b.db.QueryRow("SELECT * FROM cabify.cars WHERE ID = ?", id).Scan(&car.ID, &car.Seats)
	if err != nil {
		return types.Car{}, err
	}
	return car, nil
}

func (b *Bbdd) GetJourney(id int64) (types.JourneyDB, error) {
	var journey types.JourneyDB
	err := b.db.QueryRow("SELECT ID, people, status FROM cabify.journeys WHERE ID = ?", id).Scan(&journey.ID, &journey.People, &journey.Status)
	if err != nil {
		return types.JourneyDB{}, err
	}
	return journey, nil
}

func (b *Bbdd) UpdateJourneyStatus(journey int64) error {
	err := b.db.QueryRow("UPDATE cabify.journeys SET status = ? WHERE ID = ? ", types.Asigned, journey).Err()
	if err != nil {
		return err
	}
	return nil
}

func (b *Bbdd) GetCarByJourney(id int64) (types.Car, error) {
	var car types.Car
	err := b.db.QueryRow("SELECT ID, seats FROM cabify.cars WHERE ID_Journey = ?", id).Scan(&car.ID, &car.Seats)
	if err != nil {
		return types.Car{}, err
	}
	return car, nil
}
