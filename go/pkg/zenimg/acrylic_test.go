package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestAcrylic(t *testing.T) {
	assertEqual(t, "AC_200X200.jpg", zenimg.NewAcrylic().
		String())

	assertEqual(t, "AC_S_200X200.jpg", zenimg.NewAcrylic().
		WithStandoff().
		String())

	assertEqual(t, "AC_ED1.2_200X200.jpg", zenimg.NewAcrylic().
		EdgeDepth(1.2).
		String())

	assertEqual(t, "AC_FABCDE_200X200.jpg", zenimg.NewAcrylic().
		FrameCode("ABCDE").
		String())
}
