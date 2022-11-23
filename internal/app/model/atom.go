package model

import (
	"go-archive/pkg/db"
)

type AtomType int

//const (
//	AtomTypeNone AtomType = iota
//	AtomTypeMessage
//	AtomTypeTodo
//	AtomTypeVideo
//	AtomTypeMusic
//)

type Atom struct {
	ID      int      `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Type    AtomType `json:"type,omitempty"`
	Content string   `json:"content,omitempty"`
}

func (a *Atom) GetID() int {
	return a.ID
}

func (a *Atom) SetID(id int) {
	a.ID = id
}

func (a *Atom) AfterFind(db *db.Database) error {
	*a = Atom(*a)
	return nil
}

func QueryHosts(db *db.Database, queryFn func(a Atom) bool, limit int) ([]Atom, error) {
	var results []Atom
	var err error

	ids, err := db.IDs(AtomTableName)
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		a := Atom{}

		if err = db.Find(id, &a); err != nil {
			return nil, err
		}

		if queryFn(a) {
			results = append(results, a)
		}

		if limit != 0 && limit == len(results) {
			break
		}
	}

	return results, err
}

const AtomTableName = "atoms"

func (a *Atom) Table() string {
	return AtomTableName
}
