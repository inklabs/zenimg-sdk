package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestAluminum(t *testing.T) {
	assertEqual(t, "AL_200X200.jpg", zenimg.NewAluminum().
		String())

	assertEqual(t, "AL_ED1.2_200X200.jpg", zenimg.NewAluminum().
		EdgeDepth(1.2).
		String())

	assertEqual(t, "AL_RD3.4_200X200.jpg", zenimg.NewAluminum().
		RoundedDepth(3.4).
		String())

	assertEqual(t, "AL_FABCDE_200X200.jpg", zenimg.NewAluminum().
		FrameCode("ABCDE").
		String())
}
