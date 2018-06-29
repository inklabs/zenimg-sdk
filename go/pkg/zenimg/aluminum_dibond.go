package zenimg

import (
	"fmt"
)

type aluminumDibond struct {
	*base
	style int
}

func NewAluminumDibond() *aluminumDibond {
	p := &aluminumDibond{}
	p.base = NewBase(p)
	return p
}

func (t *aluminumDibond) Style(style int) *aluminumDibond {
	t.style = style
	return t
}

func (t *aluminumDibond) Options() []string {
	options := []string{"ALD"}

	if t.style > 0 {
		options = append(options, fmt.Sprintf("S%d", t.style))
	}

	return options
}
