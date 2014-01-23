package models

import (
  "github.com/spacedock-io/registry/db"
)

type Image struct {
  Id        int64
  Uuid      string `sql:"not null"`
  Json      []byte
  Checksum  string
  Size      int64
  Ancestry  []byte
  Tags      []Tag
}

/* type Ancestor struct { */
/*   Id int64 */
/*   ImageId int64 */
/*   /1* ParentId int64 *1/ */
/*   Value string */
/* } */

func GetImage(uuid string) (*Image, error) {
  i := &Image{}
  q := db.DB.Where("Uuid = ?", uuid).Find(i)
  if q.Error != nil {
    if q.RecordNotFound() {
      return nil, NotFoundErr
    } else {
      return nil, q.Error
    }
  }
  return i, nil
}

func SetImageChecksum(uuid string, checksum string) error {
  i, err := GetImage(uuid)
  if err != nil {
    return err
  }

  i.Checksum = checksum
  return i.Save()
}

func (image *Image) Save() error {
  q := db.DB.Save(image)
  return q.Error
}
