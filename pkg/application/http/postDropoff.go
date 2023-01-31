package application

import (
	types "github.com/isteca85/car-pooling-challenge/pkg/domain"
	"net/http"
	"strconv"
)

func (s *Server) PostDropOff(w http.ResponseWriter, r *http.Request) {
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
	if journey.Status == types.Asigned {
		err := s.DataBase.UnlinkCarsByJourney(journey.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = s.DataBase.RemoveJourney(journey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		err = s.DataBase.RemoveJourney(journey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
