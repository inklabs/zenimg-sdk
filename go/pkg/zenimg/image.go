package zenimg

import (
	"fmt"
)

type imageBuilder struct {
	imageServer string
}

func NewImageBuilder(imageServer string) *imageBuilder {
	return &imageBuilder{
		imageServer: imageServer,
	}
}

func (t imageBuilder) GetImageUrl(imageCode string, image Image) string {
	return fmt.Sprintf("%s/%s_%s", t.imageServer, imageCode, image.String())
}
