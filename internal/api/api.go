package api

import (
	"encoding/json"
	"github.com/aventhis/practice_avito/internal/auth"
	"github.com/aventhis/practice_avito/internal/models"
	"github.com/aventhis/practice_avito/internal/storage/postgres"
	"net/http"
)

type API struct {
	storage     *postgres.Storage
	authService *auth.AuthService
	router      *http.ServeMux
}

func NewAPI(storage *postgres.Storage, authService *auth.AuthService) *API {
	a := &API{
		storage:     storage,
		authService: authService,
		router:      http.NewServeMux(),
	}
	return a
}

func (a *API) SetupRoutes() http.Handler {
	a.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello, World!"))
	})

	return a.router
}

func (a *API) respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(models.ErrorResponse{Message: message})
}

func (a *API) dummyLoginHandler(writer http.ResponseWriter, request *http.Request) {
	var req models.DummyLoginRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		a.respondWithError(writer, http.StatusBadRequest, ErrInvalidJson)
		return
	}

	if !models.IsValidRole(req.Role) {
		a.respondWithError(writer, http.StatusBadRequest, ErrInvalidRole)
		return
	}

	token, err := a.authService.GenerateToken(req.Role)
	if err != nil {
		a.respondWithError(writer, http.StatusInternalServerError, ErrTokenGeneration)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(models.TokenResponse{Token: token})
}
