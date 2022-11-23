package dao

import (
	"go-archive/internal/app/model"
	"strings"
)

type Atom struct{}

func NewAtom() *Atom {
	return &Atom{}
}

func (Atom) Create(atom model.Atom) (id int, err error) {
	id, err = gdb.Insert(&atom)
	return id, err
}

func (Atom) Find(id int) (atom model.Atom, err error) {
	err = gdb.Find(id, &atom)
	return
}

func (Atom) Delete(id int) (err error) {
	err = gdb.Delete(model.AtomTableName, id)
	return
}

func (Atom) FindLikeContent(search string) (result []model.Atom, err error) {
	result = make([]model.Atom, 0)
	ids, err := gdb.IDs(model.AtomTableName)
	for _, id := range ids {
		bean := model.Atom{}
		err := gdb.Find(id, &bean)
		if err != nil {
			return nil, err
		}
		if find := strings.Contains(bean.Content, search); find {
			result = append(result, bean)
		}
	}
	return result, err
}

func (Atom) FindByType(ty model.AtomType) (result []model.Atom, err error) {
	result = make([]model.Atom, 0)
	ids, err := gdb.IDs(model.AtomTableName)
	for _, id := range ids {
		bean := model.Atom{}
		err := gdb.Find(id, &bean)
		if err != nil {
			return nil, err
		}
		if bean.Type == ty {
			result = append(result, bean)
		}
	}
	return result, err
}

func (a Atom) UpdateContent(id int, content string) (result model.Atom, err error) {
	err = gdb.Find(id, &result)
	if err != nil {
		return result, err
	}
	result.Content = content
	err = gdb.Update(&result)
	return result, err
}
