package zenimg

import "strconv"

type poster struct {
	*base
	curlDirection string
	curlWidth     string
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

func (t *poster) CurlWidth(curlWidth int) *poster {
	t.curlWidth = strconv.Itoa(curlWidth)
	return t
}

func (t *poster) Options() []string {
	options := []string{"PO"}

	if t.curlDirection != "" {
		options = append(options, t.curlDirection)
	}

	if t.curlWidth != "" {
		options = append(options, "CW"+t.curlWidth)
	}

	return options
}
