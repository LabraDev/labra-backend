package routes

import (
	"net/http"

	"labra-backend/internal/api/handlers"
)

func Oauth(mux *http.ServeMux) {
	mux.HandleFunc("/login", handlers.LoginHandler)
}
