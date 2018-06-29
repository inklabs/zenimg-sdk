package zenimg

type paper struct {
	*base
	frameCode     string
	noFrameShadow bool
	matWidth      float64
	matColor      string
}

func NewPaper() *paper {
	p := &paper{}
	p.base = NewBase(p)
	return p
}

func (t *paper) FrameCode(code string) *paper {
	t.frameCode = code
	return t
}

func (t *paper) NoFrameShadow() *paper {
	t.noFrameShadow = true
	return t
}

func (t *paper) MatWidth(width float64) *paper {
	t.matWidth = width
	return t
}

func (t *paper) MatColor(color string) *paper {
	t.matColor = color
	return t
}

func (t *paper) Options() []string {
	options := []string{"P"}

	if t.frameCode != "" {
		options = append(options, "F"+t.frameCode)

		if t.noFrameShadow {
			options = append(options, "NSHD")
		}

		if t.matWidth > 0 {
			options = append(options, "MW"+getUpToTwoDecimal(t.matWidth))
		}

		if t.matColor != "" {
			options = append(options, "MC"+t.matColor)
		}
	}

	return options
}
