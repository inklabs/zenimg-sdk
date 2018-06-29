package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestMirrorFramed(t *testing.T) {
	assertEqual(t, "MF_FABCDE_200X200.jpg", zenimg.NewMirrorFramed().
		FrameCode("ABCDE").
		String())
}
