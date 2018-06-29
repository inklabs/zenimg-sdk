package zenimg_test

import (
	"testing"

	"github.com/inklabs/zenimg-sdk/go/pkg/zenimg"
)

func TestImageBuilder(t *testing.T) {
	// given
	imageBuilder := zenimg.NewImageBuilder("http://i.zenimg.com")
	image := zenimg.NewPoster().Size(200, 200).FormatJpg()

	// when
	url := imageBuilder.GetImageUrl("e9c4937703e644ecb5abb7ceb6233c21", image)

	// then
	assertEqual(t, "http://i.zenimg.com/e9c4937703e644ecb5abb7ceb6233c21_PO_200X200.jpg", url)
}
