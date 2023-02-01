package application_test

import (
	application "github.com/isteca85/car-pooling-challenge/pkg/application/http"
	"github.com/isteca85/car-pooling-challenge/pkg/domain"
	"github.com/isteca85/car-pooling-challenge/pkg/infrastructure"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func createDBLocate() *infrastructure.MockDB {
	mockDb := &infrastructure.MockDB{}

	car := domain.CarDB{ID_Journey: 1, ID: 1, Seats: 4}
	mockDb.CarsDB = append(mockDb.CarsDB, car)

	journey := domain.JourneyDB{ID: 1, People: 4, Status: domain.Asigned}
	mockDb.JourneysDB = append(mockDb.JourneysDB, journey)
	journey2 := domain.JourneyDB{ID: 2, People: 4, Status: domain.Unasigned}
	mockDb.JourneysDB = append(mockDb.JourneysDB, journey2)

	return mockDb
}

func TestServer_PostLocate(t *testing.T) {
	t.Run("200 response post locate", func(t *testing.T) {

		serv := &application.Server{}
		serv.InitServer()
		serv.DataBase = createDBLocate()

		data := url.Values{}
		data.Set("ID", "1")
		encodedData := data.Encode()
		req, _ := http.NewRequest(http.MethodPost, "/locate", strings.NewReader(encodedData))

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

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
	t.Run("204 response post locate", func(t *testing.T) {

		serv := &application.Server{}
		serv.InitServer()
		serv.DataBase = createDBLocate()

		data := url.Values{}
		data.Set("ID", "2")
		encodedData := data.Encode()
		req, _ := http.NewRequest(http.MethodPost, "/locate", strings.NewReader(encodedData))

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		w := httptest.NewRecorder()
		serv.Router.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		got := res.StatusCode
		want := http.StatusNoContent

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("404 response post locate", func(t *testing.T) {

		serv := &application.Server{}
		serv.InitServer()
		serv.DataBase = createDBLocate()

		data := url.Values{}
		data.Set("ID", "3")
		encodedData := data.Encode()
		req, _ := http.NewRequest(http.MethodPost, "/locate", strings.NewReader(encodedData))

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		w := httptest.NewRecorder()
		serv.Router.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()
		got := res.StatusCode
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("400 response post locate", func(t *testing.T) {

		serv := &application.Server{}
		serv.InitServer()
		serv.DataBase = createDBLocate()

		data := url.Values{}
		data.Set("ID", "a")
		encodedData := data.Encode()
		req, _ := http.NewRequest(http.MethodPost, "/locate", strings.NewReader(encodedData))

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

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
