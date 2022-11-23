package api

import (
	"go-archive/pkg/chi"
	"go-archive/pkg/chi/chiutil"
)

func SetupRoutes(mux *chi.Mux) {
	apiRouter := "/api"
	chiutil.SetupResource(mux, apiRouter,
		NewNoteResource(),
		NewBackupResource(),
	)
}
