package controllers

import (
	"net/http"

	"github.com/agusdwis/blog-app/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to Blog App")
}
