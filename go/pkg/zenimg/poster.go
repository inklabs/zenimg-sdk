package zenimg

type poster struct {
	*base
	curlDirection string
	curlWidth     float64
}

func NewPoster() *poster {
	p := &poster{}
	p.base = NewBase(p)
	return p
}

func (t *poster) CurlSouthEast() *poster {
	t.curlDirection = "CSE"
	return t
}

func (t *poster) CurlSouthWest() *poster {
	t.curlDirection = "CSW"
	return t
}

func (t *poster) CurlNorthEast() *poster {
	t.curlDirection = "CNE"
	return t
}

func (t *poster) CurlNorthWest() *poster {
	t.curlDirection = "CNW"
	return t
}

func (t *poster) CurlWidth(curlWidth float64) *poster {
	t.curlWidth = curlWidth
	return t
}

func (t *poster) Options() []string {
	options := []string{"PO"}

	if t.curlDirection != "" {
		options = append(options, t.curlDirection)
	}

	if t.curlWidth > 0 {
		options = append(options, "CW"+getUpToTwoDecimal(t.curlWidth))
	}

	return options
}
