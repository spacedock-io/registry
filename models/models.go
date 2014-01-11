package models

type User struct {
  Id        int64
  Username  string
  Emails    []string
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

type Image struct {
  Id        int64
  Uuid      string
  Json      []byte
  Checksum  string
  Size      int64
  Ancestry  []string
}
