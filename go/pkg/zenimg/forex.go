package zenimg

type forex struct {
	*base
	edgeColor string
}

func NewForex() *forex {
	p := &forex{}
	p.base = NewBase(p)
	return p
}

func (t *forex) EdgeColor(color string) *forex {
	t.edgeColor = color
	return t
}

func (t *forex) Options() []string {
	options := []string{"F"}

	if t.edgeColor != "" {
		options = append(options, "EC"+t.edgeColor)
	}

	return options
}
