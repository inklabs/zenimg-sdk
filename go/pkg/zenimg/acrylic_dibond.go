package zenimg

type acrylicDibond struct {
	*base
}

func NewAcrylicDibond() *acrylicDibond {
	p := &acrylicDibond{}
	p.base = NewBase(p)
	return p
}

func (t *acrylicDibond) Options() []string {
	options := []string{"ACD"}

	return options
}
