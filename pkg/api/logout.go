package api

import (
	"net/http"
)

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	s.JSONResponse(w, r, http.StatusOK)
}

