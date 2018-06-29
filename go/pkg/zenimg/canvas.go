package zenimg

type canvas struct {
	*base
	frameCode string
}

func NewCanvas() *canvas {
	g := &canvas{}
	g.base = NewBase(g)
	return g
}

func (c *canvas) FrameCode(code string) *canvas {
	c.frameCode = code
	return c
}

func (c *canvas) Options() []string {
	options := []string{"CA"}

	if c.frameCode != "" {
		options = append(options, "F"+c.frameCode)
	}

	return options
}
