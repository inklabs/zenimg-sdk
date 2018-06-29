package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestFixedFramedPaper(t *testing.T) {
	assertEqual(t, "FP_FABCDE_S1.2X3.4_MMW5.6_200X200.jpg", zenimg.NewFixedFramedPaper().
		FrameCode("ABCDE").
		FrameSize(1.2, 3.4).
		MinMatWidth(5.6).
		String())

	assertEqual(t, "FP_FABCDE_S1.2X3.4_NSHD_MMW5.6_200X200.jpg", zenimg.NewFixedFramedPaper().
		FrameCode("ABCDE").
		FrameSize(1.2, 3.4).
		NoFrameShadow().
		MinMatWidth(5.6).
		String())

	assertEqual(t, "FP_FABCDE_S1.2X3.4_MMW5.6_MC00F_200X200.jpg", zenimg.NewFixedFramedPaper().
		FrameCode("ABCDE").
		FrameSize(1.2, 3.4).
		MinMatWidth(5.6).
		MatColor("00F").
		String())
}
