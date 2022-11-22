package bind

import (
	"go-archive/internal/app/model"
	"net/http"
)

type AtomBind struct {
	model.Atom
}

func (a *AtomBind) Model() *model.Atom {
	return &a.Atom
}

func (a *AtomBind) Bind(r *http.Request) error {
	return nil
}
