package zenimg

type aluminum struct {
	*base
	edgeDepth    float64
	roundedDepth float64
	frameCode    string
}

func NewAluminum() *aluminum {
	p := &aluminum{}
	p.base = NewBase(p)
	return p
}

func (t *aluminum) EdgeDepth(depth float64) *aluminum {
	t.edgeDepth = depth
	return t
}

func (t *aluminum) RoundedDepth(depth float64) *aluminum {
	t.roundedDepth = depth
	return t
}

func (t *aluminum) FrameCode(code string) *aluminum {
	t.frameCode = code
	return t
}

func (t *aluminum) Options() []string {
	options := []string{"AL"}

	if t.edgeDepth > 0 {
		options = append(options, "ED"+getUpToTwoDecimal(t.edgeDepth))
	}

	if t.roundedDepth > 0 {
		options = append(options, "RD"+getUpToTwoDecimal(t.roundedDepth))
	}

	if t.frameCode != "" {
		options = append(options, "F"+t.frameCode)
	}

	return options
}
