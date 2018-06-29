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
	actualWidth            float64
	actualHeight           float64
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

func (t *base) ActualSize(width float64, height float64) *base {
	t.actualWidth = width
	t.actualHeight = height
	return t
}

func (t *base) Size(width int, height int) *base {
	t.width = width
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
		pieces = append(pieces, "A"+formatSize(t.actualWidth, t.actualHeight))
	}

	pieces = append(pieces, fmt.Sprintf("%dX%d", t.width, t.height))

	return strings.Join(pieces, `_`) + `.` + t.format
}

func formatSize(width float64, height float64) string {
	return fmt.Sprintf("%sX%s",
		getUpToTwoDecimal(width),
		getUpToTwoDecimal(height),
	)
}

func getUpToTwoDecimal(val float64) string {
	return strings.TrimRight(
		strings.TrimRight(fmt.Sprintf("%.2f", val), "0"),
		".",
	)
}
