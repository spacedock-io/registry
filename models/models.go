package models

import(
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

var DB gorm.DB

func init() {
  var err error
  DB, err = gorm.Open("postgresql", "user=yawnt dbname=test sslmode=disable")

  if err != nil {
    panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
  }
}

type Email  string

type User struct {
  Id        int64
  Username  string
  Emails    []Email
  Hash      []byte
  Salt      []byte
}

type Token struct {
  Id        int64
  Signature string
  Access    string
}

type Repo struct {
  Id        int64
  Name      string
  Owner     User
  Tokens    []Token
}

type UUID   string

type Image struct {
  Id        int64
  Uuid      UUID
  Json      string
  Checksum  string
  Size      int64
  Ancestry  []UUID
}
