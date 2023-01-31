package application

import (
	"encoding/json"
	"github.com/isteca85/car-pooling-challenge/pkg/domain"
	"net/http"
	"strconv"
)

func (s *Server) PostLocate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("ID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	journey, err := s.DataBase.GetJourney(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if journey.Status == domain.Asigned {
		w.WriteHeader(http.StatusOK)
		car, err := s.DataBase.GetCarByJourney(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		result, err := json.Marshal(car)
		w.Write([]byte(result))
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
