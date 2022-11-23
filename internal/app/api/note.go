package api

import (
	"errors"
	"go-archive/internal/app/dao"
	"go-archive/internal/app/service"
	"go-archive/internal/pkg/bind"
	"go-archive/pkg/chi"
	"go-archive/pkg/chi/chiutil"
	"go-archive/pkg/chi/render"
	"go-archive/pkg/chi/response"
	"net/http"
	"strconv"
)

type NoteResource struct {
	dAtom *dao.Atom
	sAtom *service.Atom
}

func NewNoteResource() chiutil.Resource {
	return &NoteResource{
		dAtom: dao.NewAtom(),
		sAtom: service.NewAtom(),
	}
}

func (p *NoteResource) Register(router chi.Router) {
	router.Get("/note", p.findAll)
	router.Get("/note/{id}", p.find)
	router.Post("/note", p.create)
	router.Delete("/note/{id}", p.delete)
	router.Put("/note/{id}/content", p.update)
}

func (p *NoteResource) findAll(w http.ResponseWriter, r *http.Request) {
	atomList, err := p.sAtom.List()
	if err != nil {
		response.JSONServerError(w, err)
		return
	}
	response.JSONList(w, atomList, int64(len(atomList)))
}

func (p *NoteResource) find(w http.ResponseWriter, r *http.Request) {
	noteID := chi.URLParam(r, "id")
	if noteID == "" {
		response.JSONBadRequest(w, errors.New("找不到对于记录"))
		return
	}
	id, err := strconv.Atoi(noteID)
	if err != nil {
		response.JSONBadRequest(w, err)
		return
	}
	atom, err := p.dAtom.Find(id)
	if err != nil {
		response.JSONBadRequest(w, err)
		return
	}
	response.JSONData(w, atom)
}

func (p *NoteResource) create(w http.ResponseWriter, r *http.Request) {
	data := bind.AtomBind{}
	if err := render.Bind(r, &data); err != nil {
		response.JSONBadRequest(w, err)
		return
	}
	_, err := p.sAtom.Add(data.Model())
	if err != nil {
		response.JSONServerError(w, err)
		return
	}
	response.JSONData(w, data)
}

func (p *NoteResource) update(w http.ResponseWriter, r *http.Request) {
	noteID := chi.URLParam(r, "id")
	if noteID == "" {
		response.JSONBadRequest(w, errors.New("找不到对于记录"))
		return
	}
	id, err := strconv.Atoi(noteID)
	if err != nil {
		response.JSONBadRequest(w, err)
		return
	}
	data := bind.AtomBind{}
	if err := render.Bind(r, &data); err != nil {
		response.JSONBadRequest(w, err)
		return
	}
	result, err := p.dAtom.UpdateContent(id, data.Content)
	if err != nil {
		response.JSONServerError(w, err)
		return
	}
	response.JSONData(w, result)
}

func (p *NoteResource) delete(w http.ResponseWriter, r *http.Request) {
	if id, err := chi.URLParamInt(r, "id"); err != nil {
		response.JSONBadRequest(w, err)
		return
	} else {
		p.dAtom.Delete(id)
	}
	response.JSON(w)
}
