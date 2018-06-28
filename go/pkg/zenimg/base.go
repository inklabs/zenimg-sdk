package zenimg

import (
	"fmt"
	"strconv"
	"strings"
)

type Substrate interface {
	Options() []string
}

type Image interface {
	String() string
}

type axis struct {
	pan  int16
	tilt int16
	roll int16
}

type base struct {
	background             string
	backgroundTexture      string
	backgroundTextureColor string
	shadow                 string
	axis                   axis
	actualWidth            int16
	actualHeight           int16
	width                  int
	height                 int
	format                 string
	substrate              Substrate
}

func NewBase(substrate Substrate) *base {
	return &base{
		width:     200,
		height:    200,
		format:    "jpg",
		substrate: substrate,
	}
}

func (t *base) EnableShadow() *base {
	t.shadow = "SHD"
	return t
}

func (t *base) Background(i int) *base {
	t.background = "BG" + strconv.Itoa(i)
	return t
}

func (t *base) BackgroundTexture(i int) *base {
	t.backgroundTexture = "BT" + strconv.Itoa(i)
	return t
}

func (t *base) BackgroundTextureColor(hexColor string) *base {
	t.backgroundTextureColor = hexColor
	return t
}

func (t *base) Pan(i int16) *base {
	t.axis.pan = i
	return t
}

func (t *base) Tilt(i int16) *base {
	t.axis.tilt = i
	return t
}

func (t *base) Roll(i int16) *base {
	t.axis.roll = i
	return t
}

func (t *base) ActualWidth(width int16) *base {
	t.actualWidth = width
	return t
}

func (t *base) ActualHeight(width int16) *base {
	t.actualHeight = width
	return t
}

func (t *base) Width(width int) *base {
	t.width = width
	return t
}

func (t *base) Height(height int) *base {
	t.height = height
	return t
}

func (t *base) FormatPng() *base {
	t.format = "png"
	return t
}

func (t *base) FormatJpg() *base {
	t.format = "jpg"
	return t
}

func (t *base) String() string {
	pieces := t.substrate.Options()

	if t.background != "" {
		pieces = append(pieces, t.background)
	} else if t.backgroundTexture != "" {
		pieces = append(pieces, t.backgroundTexture)

		if t.backgroundTextureColor != "" {
			pieces = append(pieces, "BTC"+t.backgroundTextureColor)
		}
	} else if t.shadow != "" {
		pieces = append(pieces, t.shadow)
	}

	if t.axis.pan != 0 {
		pieces = append(pieces, fmt.Sprintf("P%d", t.axis.pan))
	}

	if t.axis.tilt != 0 {
		pieces = append(pieces, fmt.Sprintf("T%d", t.axis.tilt))
	}

	if t.axis.roll != 0 {
		pieces = append(pieces, fmt.Sprintf("R%d", t.axis.roll))
	}

	if t.actualWidth != 0 && t.actualHeight != 0 {
		pieces = append(pieces, fmt.Sprintf("A%dX%d", t.actualWidth, t.actualHeight))
	}

	pieces = append(pieces, fmt.Sprintf("%dX%d", t.width, t.height))

	return strings.Join(pieces, `_`) + `.` + t.format
}
