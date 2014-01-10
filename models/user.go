package models

import(
  "github.com/jinzhu/gorm"
)

type Email string

type User struct {
  Id       int64
  Username string
  Emails   []Email
  Hash     []byte
  Salt     []byte
}
