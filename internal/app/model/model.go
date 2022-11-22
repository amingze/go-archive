package model

func Tables() []Register {
	return []Register{
		new(Atom),
	}
}

type Register interface {
	Table() string
}
