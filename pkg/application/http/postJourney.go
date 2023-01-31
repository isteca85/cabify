package application

import (
	"encoding/json"
	"fmt"
	types "github.com/isteca85/car-pooling-challenge/pkg/domain"
	"io"
	"log"
	"net/http"
)

func (s *Server) PostJourney(w http.ResponseWriter, r *http.Request) {
	var journey types.Journey
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err = json.Unmarshal(reqBody, &journey)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			_, err := s.DataBase.InsertJourney(journey)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			freeCar, err := s.DataBase.GetFreeCar(journey.People)
			if freeCar > 0 {
				err := s.DataBase.UpdateCarByJourney(freeCar, journey.ID)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusAccepted)
			}
		}
	}
}
