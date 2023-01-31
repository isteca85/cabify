package application

import (
	"github.com/gorilla/mux"
	infra "github.com/isteca85/car-pooling-challenge/pkg/infrastructure"
	"net/http"
)

type Server struct {
	Router   http.Handler
	DataBase infra.DataBase
}

func (s *Server) InitServer() {
	r := mux.NewRouter()
	r.HandleFunc("/status", s.GetStatus).Methods(http.MethodGet)
	r.HandleFunc("/cars", s.PutCars).Methods(http.MethodPut)
	r.HandleFunc("/dropoff", s.PostDropOff).Methods(http.MethodPost)
	r.HandleFunc("/journey", s.PostJourney).Methods(http.MethodPost)
	r.HandleFunc("/locate", s.PostLocate).Methods(http.MethodPost)
	s.Router = r
}
