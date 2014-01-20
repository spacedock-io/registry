package models

type Image struct {
  Id        int64
  Uuid      string
  Json      []byte
  Checksum  string
  Size      int64
  Ancestry  []string
  Tags      []Tag
}
