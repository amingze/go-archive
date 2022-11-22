package service

import (
	"go-archive/internal/app/dao"
	"go-archive/internal/app/model"
)

type Atom struct {
	dAtom *dao.Atom
}

func NewAtom() *Atom {
	return &Atom{dao.NewAtom()}
}

func (s Atom) Add(atom *model.Atom) (int, error) {
	id, err := s.dAtom.Create(*atom)
	return id, err
}

func (s Atom) List() ([]model.Atom, error) {
	atoms, err := s.dAtom.FindLikeContent("")
	return atoms, err
}
