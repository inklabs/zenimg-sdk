package zenimg_test

import (
	"testing"

	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
)

func TestWood(t *testing.T) {
	assertEqual(t, "WD_200X200.jpg", zenimg.NewWood().
		String())

	assertEqual(t, "WD_MA_200X200.jpg", zenimg.NewWood().
		Maple().
		String())

	assertEqual(t, "WD_BD_200X200.jpg", zenimg.NewWood().
		BambooDark().
		String())

	assertEqual(t, "WD_BL_200X200.jpg", zenimg.NewWood().
		BambooLight().
		String())

	assertEqual(t, "WD_BE_200X200.jpg", zenimg.NewWood().
		Beech().
		String())

	assertEqual(t, "WD_CD_200X200.jpg", zenimg.NewWood().
		CherryDark().
		String())

	assertEqual(t, "WD_CL_200X200.jpg", zenimg.NewWood().
		CherryLight().
		String())

	assertEqual(t, "WD_OK_200X200.jpg", zenimg.NewWood().
		Oak().
		String())

	assertEqual(t, "WD_WN_200X200.jpg", zenimg.NewWood().
		Walnut().
		String())

	assertEqual(t, "WD_WN_200X200.jpg", zenimg.NewWood().
		Walnut().
		FormatJpg().
		String())

	assertEqual(t, "WD_WN_P15_T-25_R35_200X200.jpg", zenimg.NewWood().
		Walnut().
		Pan(15).
		Tilt(-25).
		Roll(35).
		String())

	assertEqual(t, "WD_WN_A10X10_200X200.jpg", zenimg.NewWood().
		Walnut().
		ActualWidth(10).
		ActualHeight(10).
		String())
}
