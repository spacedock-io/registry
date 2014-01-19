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
  Ancestry  []string
  Tags      []Tag
}

type Tag struct {
  Id         int64
  Tag        string `sql:"not null"`
  ImageId    int64
  Repo       string `sql:"not null"`
  Namespace  string `sql:"not null"`
}

func GetTags(namespace string, repo string) ([]Tag, error) {
  var t []Tag
  q := db.DB.Where("Namespace = ? and Repo = ?", namespace, repo).Find(&t)
  if q.Error != nil {
    return nil, q.Error
  }
  return t, nil
}

func (tag *Tag) Save() error {
  q := db.DB.Save(tag)
  return q.Error
}

func (tag *Tag) Create() error {
  return tag.Save()
}
