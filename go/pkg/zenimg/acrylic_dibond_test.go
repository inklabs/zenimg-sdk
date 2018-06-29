package zenimg_test

import (
	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
	"testing"
)

func TestAcrylicDibond(t *testing.T) {
	assertEqual(t, "ACD_200X200.jpg", zenimg.NewAcrylicDibond().
		String())
}
