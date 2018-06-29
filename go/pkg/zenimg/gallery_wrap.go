package zenimg

type galleryWrap struct {
	*base
	withImageWrap bool
	edgeColor     string
	depth         float64
}

func NewGalleryWrap() *galleryWrap {
	g := &galleryWrap{}
	g.base = NewBase(g)
	return g
}

func (g *galleryWrap) WithImageWrap() *galleryWrap {
	g.withImageWrap = true
	return g
}

func (g *galleryWrap) EdgeColor(color string) *galleryWrap {
	g.edgeColor = color
	return g
}

func (g *galleryWrap) Depth(depth float64) *galleryWrap {
	g.depth = depth
	return g
}

func (g *galleryWrap) Options() []string {
	options := []string{"CG"}

	if g.withImageWrap {
		options = append(options, "IW")
	} else if g.edgeColor != "" {
		options = append(options, "EC"+g.edgeColor)
	}

	if g.depth > 0 {
		options = append(options, "D"+getUpToTwoDecimal(g.depth))
	}

	return options
}
