package zenimg_test

import (
	"testing"

	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
)

func TestPoster(t *testing.T) {
	assertEqual(t, "PO_200X200.jpg", zenimg.NewPoster().
		String())

	assertEqual(t, "PO_CSE_200X200.jpg", zenimg.NewPoster().
		CurlSouthEast().
		String())

	assertEqual(t, "PO_CW6_200X200.jpg", zenimg.NewPoster().
		CurlWidth(6).
		String())

	assertEqual(t, "PO_CSW_CW5_SHD_200X200.jpg", zenimg.NewPoster().
		CurlSouthWest().
		CurlWidth(5).
		EnableShadow().
		String())

	assertEqual(t, "PO_CNE_CW4_BG0_200X200.jpg", zenimg.NewPoster().
		CurlNorthEast().
		CurlWidth(4).
		Background(0).
		String())

	assertEqual(t, "PO_CNW_CW3_BT1_200X200.png", zenimg.NewPoster().
		CurlNorthWest().
		CurlWidth(3).
		BackgroundTexture(1).
		FormatPng().
		String())

	assertEqual(t, "PO_CNW_CW3_BT1_BTCF00_200X200.png", zenimg.NewPoster().
		CurlNorthWest().
		CurlWidth(3).
		BackgroundTexture(1).
		BackgroundTextureColor("F00").
		FormatPng().
		String())

	assertEqual(t, "PO_CNW_CW3_BG0_500X400.png", zenimg.NewPoster().
		CurlNorthWest().
		CurlWidth(3).
		Background(0).
		Width(500).
		Height(400).
		FormatPng().
		String())
}
