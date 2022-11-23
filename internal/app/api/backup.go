package api

import (
	"go-archive/internal/app/service"
	"go-archive/pkg/chi"
	"go-archive/pkg/chi/chiutil"
	"net/http"
)

type BackupResource struct {
	sBackup *service.Backup
}

func NewBackupResource() chiutil.Resource {
	return &BackupResource{
		sBackup: service.NewBackup(),
	}
}

func (p *BackupResource) Register(router chi.Router) {
	router.Post("/backup", p.create)
	router.Get("/backup", p.findAll)
}

func (p *BackupResource) create(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (p *BackupResource) findAll(w http.ResponseWriter, r *http.Request) {
	// TODO
}
