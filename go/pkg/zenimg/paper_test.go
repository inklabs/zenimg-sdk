package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestPaper(t *testing.T) {
	assertEqual(t, "P_200X200.jpg", zenimg.NewPaper().
		String())

	assertEqual(t, "P_FABCDE_200X200.jpg", zenimg.NewPaper().
		FrameCode("ABCDE").
		String())

	assertEqual(t, "P_FABCDE_NSHD_200X200.jpg", zenimg.NewPaper().
		FrameCode("ABCDE").
		NoFrameShadow().
		String())

	assertEqual(t, "P_FABCDE_MW1.23_200X200.jpg", zenimg.NewPaper().
		FrameCode("ABCDE").
		MatWidth(1.23).
		String())

	assertEqual(t, "P_FABCDE_MW1.23_MCFFF_200X200.jpg", zenimg.NewPaper().
		FrameCode("ABCDE").
		MatWidth(1.23).
		MatColor("FFF").
		String())
}
