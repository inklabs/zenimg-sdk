package zenimg

type fixedFramedPaper struct {
	*base
	frameCode     string
	frameWidth    float64
	frameHeight   float64
	noFrameShadow bool
	minMatWidth   float64
	matColor      string
}

func NewFixedFramedPaper() *fixedFramedPaper {
	g := &fixedFramedPaper{}
	g.base = NewBase(g)
	return g
}

func (c *fixedFramedPaper) FrameCode(code string) *fixedFramedPaper {
	c.frameCode = code
	return c
}

func (c *fixedFramedPaper) FrameSize(width float64, height float64) *fixedFramedPaper {
	c.frameWidth = width
	c.frameHeight = height
	return c
}

func (c *fixedFramedPaper) NoFrameShadow() *fixedFramedPaper {
	c.noFrameShadow = true
	return c
}

func (c *fixedFramedPaper) MinMatWidth(width float64) *fixedFramedPaper {
	c.minMatWidth = width
	return c
}

func (c *fixedFramedPaper) MatColor(color string) *fixedFramedPaper {
	c.matColor = color
	return c
}

func (c *fixedFramedPaper) Options() []string {
	options := []string{"FP"}

	options = append(options,
		"F"+c.frameCode,
		"S"+formatSize(c.frameWidth, c.frameHeight),
	)

	if c.noFrameShadow {
		options = append(options, "NSHD")
	}

	options = append(options, "MMW"+getUpToTwoDecimal(c.minMatWidth))

	if c.matColor != "" {
		options = append(options, "MC"+c.matColor)
	}

	return options
}
