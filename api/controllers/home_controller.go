package controllers

import (
	"net/http"

	"github.com/elieserguterres/golang2/api/responses"
)

// Home listando response
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
