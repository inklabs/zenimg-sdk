package zenimg

import (
	"strings"
	"strconv"
)

type base struct {
	background string
	shadow     string
	width      int
	height     int
	format     string
}

func NewSubstrate() *base {
	return &base{
		width:  200,
		height: 200,
		format: `jpg`,
	}
}

func (t *base) enableShadow() {
	t.shadow = `SHD`
}

func (t *base) setBackground(i int) *base {
	t.background = `BG` + strconv.Itoa(i)
	return t
}

func (t *base) setFormat(format string) *base {
	t.format = format
	return t
}

func (t *base) String(options []string) string {
	pieces := options

	if t.background != `` {
		pieces = append(pieces, t.background)
	} else if t.shadow != `` {
		pieces = append(pieces, t.shadow)
	}

	pieces = append(pieces, strconv.Itoa(t.width) + `X` + strconv.Itoa(t.height))

	return strings.Join(pieces, `_`) + `.` + t.format
}
