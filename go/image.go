package zenimg

type imageBuilder struct {
	imageServer string
}

func NewImageBuilder(imageServer string) *imageBuilder {
	return &imageBuilder{
		imageServer: imageServer,
	}
}

func (t imageBuilder) GetImageUrl(imageCode string, substrate Substrate) string {
	return t.imageServer + imageCode + `_` + substrate.String()
}
