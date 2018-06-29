package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestGalleryWrap(t *testing.T) {
	assertEqual(t, "CG_200X200.jpg", zenimg.NewGalleryWrap().
		String())

	assertEqual(t, "CG_IW_200X200.jpg", zenimg.NewGalleryWrap().
		WithImageWrap().
		String())

	assertEqual(t, "CG_EC0F0_200X200.jpg", zenimg.NewGalleryWrap().
		EdgeColor("0F0").
		String())

	assertEqual(t, "CG_D3.4_200X200.jpg", zenimg.NewGalleryWrap().
		Depth(3.4).
		String())

	assertEqual(t, "CG_D3.12_200X200.jpg", zenimg.NewGalleryWrap().
		Depth(3.1234).
		String())

	assertEqual(t, "CG_D3.13_200X200.jpg", zenimg.NewGalleryWrap().
		Depth(3.1299).
		String())

	assertEqual(t, "CG_D3_200X200.jpg", zenimg.NewGalleryWrap().
		Depth(3.0).
		String())
}
