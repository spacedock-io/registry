package cloudfiles

import (
  "github.com/ncw/swift"
  "github.com/stretchr/objx"
)

var Cloudfiles swift.Connection

func New(c objx.Map) swift.Connection {
  Cloudfiles = swift.Connection{
    UserName: c.Get("cloudfiles.username").Str(),
    ApiKey:   c.Get("cloudfiles.apiKey").Str(),
    AuthUrl:  c.Get("cloudfiles.authUrl").Str(),
  }

  err := Cloudfiles.Authenticate()
  if err != nil {
    panic(err)
  }

  return Cloudfiles
}
