package application

import (
	"encoding/json"
	"fmt"
	types "github.com/isteca85/car-pooling-challenge/pkg/domain"
	"io"
	"net/http"
)

func (s *Server) PutCars(w http.ResponseWriter, r *http.Request) {
	var cars []types.Car
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err = json.Unmarshal(reqBody, &cars)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			s.DataBase.CleanDdbb()
			for _, car := range cars {
				_, err := s.DataBase.InsertCar(car)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					s.DataBase.CleanDdbb()
					return
				}
			}
			w.WriteHeader(http.StatusOK)
		}
	}
}
