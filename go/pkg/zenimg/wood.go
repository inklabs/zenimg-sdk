package zenimg

type wood struct {
	*base
	code string
}

func NewWood() *wood {
	w := &wood{}
	w.base = NewBase(w)
	return w
}

func (t *wood) Maple() *wood {
	t.code = "MA"
	return t
}

func (t *wood) BambooDark() *wood {
	t.code = "BD"
	return t
}

func (t *wood) BambooLight() *wood {
	t.code = "BL"
	return t
}

func (t *wood) Beech() *wood {
	t.code = "BE"
	return t
}

func (t *wood) CherryDark() *wood {
	t.code = "CD"
	return t
}

func (t *wood) CherryLight() *wood {
	t.code = "CL"
	return t
}

func (t *wood) Oak() *wood {
	t.code = "OK"
	return t
}

func (t *wood) Walnut() *wood {
	t.code = "WN"
	return t
}

func (t *wood) Options() []string {
	options := []string{"WD"}

	if t.code != "" {
		options = append(options, t.code)
	}

	return options
}
