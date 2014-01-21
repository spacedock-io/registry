package models

import (
  "github.com/spacedock-io/registry/db"
)

type Tag struct {
  Id         int64
  Tag        string `sql:"not null"`
  ImageId    int64
  Repo       string `sql:"not null"`
  Namespace  string `sql:"not null"`
}

func CreateTag(namespace string, repo string, tag string, uuid string) error {
  image := &Image{}
  q := db.DB.Where("Uuid = ?", uuid).Find(image)

  if q.RecordNotFound() {
    t := Tag{
      Tag: tag,
      Repo: repo,
      Namespace: namespace,
    }
    image.Tags = append(image.Tags, t)
    return image.Save()
  } else if q.Error != nil {
    return q.Error
  }

  return TagCreateErr
}

func GetTags(namespace string, repo string) ([]Tag, error) {
  var t []Tag
  q := db.DB.Where("Namespace = ? and Repo = ?", namespace, repo).Find(&t)
  if q.Error != nil {
    return nil, q.Error
  }
  return t, nil
}

func GetTag(namespace string, repo string, tag string) (*Tag, error) {
  t := &Tag{}
  q := db.DB.Where("Namespace = ? and Repo = ?", namespace, repo).Find(t)
  if q.Error != nil {
    return nil, q.Error
  }
  return t, nil
}

func (tag *Tag) Save() error {
  q := db.DB.Save(tag)
  if q.Error != nil {
    return TagSaveErr
  }
  return nil
}

func (tag *Tag) Create() error {
  return tag.Save()
}
