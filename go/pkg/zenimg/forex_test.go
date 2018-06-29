package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestForex(t *testing.T) {
	assertEqual(t, "F_200X200.jpg", zenimg.NewForex().
		String())

	assertEqual(t, "F_ECFF0_200X200.jpg", zenimg.NewForex().
		EdgeColor("FF0").
		String())
}
