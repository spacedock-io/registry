package db

import (
  "fmt"
  "github.com/jinzhu/gorm"
  "github.com/stretchr/objx"
  _ "github.com/lib/pq"
)

var DB gorm.DB

func New(c objx.Map) gorm.DB {
  var err error
  DB, err = gorm.Open(
    c.Get("gorm.driver").Str(),
    c.Get("gorm.source").Str())

  if err != nil {
    panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
  }

  return DB
}
