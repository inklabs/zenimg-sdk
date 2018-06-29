package zenimg

type mirrorFramed struct {
	*base
	frameCode string
}

func NewMirrorFramed() *mirrorFramed {
	p := &mirrorFramed{}
	p.base = NewBase(p)
	return p
}

func (t *mirrorFramed) FrameCode(code string) *mirrorFramed {
	t.frameCode = code
	return t
}

func (t *mirrorFramed) Options() []string {
	options := []string{"MF", "F" + t.frameCode}

	return options
}
