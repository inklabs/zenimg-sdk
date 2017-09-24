package zenimg

import "testing"

func TestPoster(t *testing.T) {
	assertEqual(t, `PO_200X200.jpg`, NewPoster().
		String())

	assertEqual(t, `PO_CSE_200X200.jpg`, NewPoster().
		CurlSouthEast().
		String())

	assertEqual(t, `PO_CW6_200X200.jpg`, NewPoster().
		CurlWidth(6).
		String())

	assertEqual(t, `PO_CSW_CW5_SHD_200X200.jpg`, NewPoster().
		CurlSouthWest().
		CurlWidth(5).
		EnableShadow().
		String())

	assertEqual(t, `PO_CNE_CW4_BG0_200X200.jpg`, NewPoster().
		CurlNorthEast().
		CurlWidth(4).
		Background(0).
		String())

	assertEqual(t, `PO_CNW_CW3_BG0_200X200.png`, NewPoster().
		CurlNorthWest().
		CurlWidth(3).
		Background(0).
		FormatPng().
		String())

	assertEqual(t, `PO_CNW_CW3_BG0_500X400.png`, NewPoster().
		CurlNorthWest().
		CurlWidth(3).
		Background(0).
		Width(500).Height(400).
		FormatPng().
		String())
}

func assertEqual(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Error(`Expected [` + expected + `] got [` + actual + `]`)
	}
}
