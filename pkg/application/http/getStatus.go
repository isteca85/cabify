package application

import "net/http"

func (s *Server) GetStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
