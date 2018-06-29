package zenimg

type acrylic struct {
	*base
	withStandoff bool
	edgeDepth    float64
	frameCode    string
}

func NewAcrylic() *acrylic {
	p := &acrylic{}
	p.base = NewBase(p)
	return p
}

func (t *acrylic) WithStandoff() *acrylic {
	t.withStandoff = true
	return t
}

func (t *acrylic) EdgeDepth(depth float64) *acrylic {
	t.edgeDepth = depth
	return t
}

func (t *acrylic) FrameCode(code string) *acrylic {
	t.frameCode = code
	return t
}

func (t *acrylic) Options() []string {
	options := []string{"AC"}

	if t.withStandoff {
		options = append(options, "S")
	}

	if t.edgeDepth > 0 {
		options = append(options, "ED"+getUpToTwoDecimal(t.edgeDepth))
	}

	if t.frameCode != "" {
		options = append(options, "F"+t.frameCode)
	}

	return options
}
