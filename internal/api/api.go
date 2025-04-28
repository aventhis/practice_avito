package api

import (
	"github.com/aventhis/practice_avito/internal/auth"
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
