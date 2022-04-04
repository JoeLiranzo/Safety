package controllers

import (
	"net/http"

	"deepthinking.do/safetygo/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To SafetyGO API")
}
