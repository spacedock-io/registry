package models

import (
  "github.com/spacedock-io/registry/db"
)

type Image struct {
  Id        int64
  Uuid      string
  Json      []byte
  Checksum  string
  Size      int64
  Ancestry  []Ancestor
  Tags      []Tag
}

type Ancestor struct {
  Id int64
  ImageId int64
  Name string
}

func (image *Image) Save() error {
  q := db.DB.Save(image)
  return q.Error
}
