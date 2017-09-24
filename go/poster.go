package zenimg

import "strconv"

type poster struct {
	base          *base
	curlDirection string
	curlWidth     string
}

func NewPoster() *poster {
	return &poster{base: NewSubstrate()}
}

func (t *poster) CurlSouthEast() *poster {
	t.curlDirection = `CSE`
	return t
}

func (t *poster) CurlSouthWest() *poster {
	t.curlDirection = `CSW`
	return t
}

func (t *poster) CurlNorthEast() *poster {
	t.curlDirection = `CNE`
	return t
}

func (t *poster) CurlNorthWest() *poster {
	t.curlDirection = `CNW`
	return t
}

func (t *poster) CurlWidth(curlWidth int) *poster {
	t.curlWidth = strconv.Itoa(curlWidth)
	return t
}

func (t *poster) EnableShadow() *poster {
	t.base.enableShadow()
	return t
}

func (t *poster) Background(i int) *poster {
	t.base.setBackground(i)
	return t
}

func (t *poster) Width(width int) *poster {
	t.base.width = width
	return t
}

func (t *poster) Height(height int) *poster {
	t.base.height = height
	return t
}

func (t *poster) FormatPng() *poster {
	t.base.setFormat(`png`)
	return t
}

func (t *poster) FormatJpg() *poster {
	t.base.setFormat(`jpg`)
	return t
}

func (t *poster) String() string {
	options := []string{`PO`}

	if t.curlDirection != `` {
		options = append(options, t.curlDirection)
	}

	if t.curlWidth != `` {
		options = append(options, `CW` + t.curlWidth)
	}

	return t.base.String(options)
}
