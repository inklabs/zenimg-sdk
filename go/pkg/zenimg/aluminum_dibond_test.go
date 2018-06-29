package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestAluminumDibond(t *testing.T) {
	assertEqual(t, "ALD_200X200.jpg", zenimg.NewAluminumDibond().
		String())

	assertEqual(t, "ALD_S1_200X200.jpg", zenimg.NewAluminumDibond().
		Style(1).
		String())
}
