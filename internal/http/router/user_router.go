package router

import (
	"net/http"

	"github.com/medorahealth/medora_backend/internal/http/handler"

	"github.com/go-chi/chi/v5"
)

func UserRouter(userHandler *handler.UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", userHandler.RegisterUserHandler)
	r.Post("/login", userHandler.LoginUserHandler)

	return r
}
