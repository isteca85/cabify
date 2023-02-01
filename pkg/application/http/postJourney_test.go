package application_test

import (
	"bytes"
	application "github.com/isteca85/car-pooling-challenge/pkg/application/http"
	"github.com/isteca85/car-pooling-challenge/pkg/domain"
	"github.com/isteca85/car-pooling-challenge/pkg/infrastructure"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createDBJourney() *infrastructure.MockDB {
	mockDb := &infrastructure.MockDB{}

	car := domain.CarDB{ID_Journey: 0, ID: 1, Seats: 4}
	mockDb.CarsDB = append(mockDb.CarsDB, car)
	journey := domain.JourneyDB{ID: 2, People: 5, Status: domain.Unasigned}
	mockDb.JourneysDB = append(mockDb.JourneysDB, journey)

	return mockDb
}

func TestServer_PostJourney(t *testing.T) {
	t.Run("200 response post journey", func(t *testing.T) {
		serv := &application.Server{}
		serv.InitServer()
		serv.DataBase = createDBJourney()

		var jsonStr = []byte(`{"ID": 1, "People": 4}`)
		req, _ := http.NewRequest(http.MethodPost, "/journey", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		serv.Router.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		got := res.StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("202 response post journey", func(t *testing.T) {
		serv := &application.Server{}
		serv.InitServer()
		serv.DataBase = createDBJourney()

		var jsonStr = []byte(`{"ID": 1, "People": 6}`)
		req, _ := http.NewRequest(http.MethodPost, "/journey", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		serv.Router.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		got := res.StatusCode
		want := http.StatusAccepted

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("400 response post journey", func(t *testing.T) {
		serv := &application.Server{}
		serv.InitServer()
		serv.DataBase = createDBJourney()

		var jsonStr = []byte(`{"ID": 1, "People": A}`)
		req, _ := http.NewRequest(http.MethodPost, "/journey", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		serv.Router.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		got := res.StatusCode
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
