package application_test

import (
	application "github.com/isteca85/car-pooling-challenge/pkg/application/http"
	"github.com/isteca85/car-pooling-challenge/pkg/infrastructure"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_GetStatus(t *testing.T) {
	t.Run("succesfull response", func(t *testing.T) {

		mockDb := &infrastructure.MockDB{}
		serv := &application.Server{}
		serv.InitServer()
		serv.DataBase = mockDb

		req := httptest.NewRequest(http.MethodGet, "/status", nil)
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
}
