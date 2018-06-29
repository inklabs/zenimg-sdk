package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestCanvas(t *testing.T) {
	assertEqual(t, "CA_200X200.jpg", zenimg.NewCanvas().
		String())

	assertEqual(t, "CA_FABCDE_200X200.jpg", zenimg.NewCanvas().
		FrameCode("ABCDE").
		String())
}
