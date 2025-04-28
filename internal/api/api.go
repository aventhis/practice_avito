package api

import (
	"github.com/aventhis/practice_avito/internal/auth"
	"github.com/aventhis/practice_avito/internal/storage/postgres"
)

type API struct {
	storage     *postgres.Storage
	authService *auth.AuthService
}

func NewAPI(storage *postgres.Storage, authService *auth.AuthService) *API {
	return &API{storage: storage, authService: authService}
}
