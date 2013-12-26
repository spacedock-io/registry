package images

const IMAGE_BASE_PATH = "images/"

type Image struct {
  id string
}

func (i *Image) JsonPath() string {
  return IMAGE_BASE_PATH + i.id  + "/json"
}

func (i *Image) LayerPath() string {
  return IMAGE_BASE_PATH + i.id  + "/layer"
}

func (i *Image) AncestryPath() string {
  return IMAGE_BASE_PATH + i.id + "/ancestry"
}

