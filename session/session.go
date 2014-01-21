package sx

import (
  "net/http"
  "github.com/gorilla/sessions"
  "github.com/ricallinson/stackr"
  "github.com/ricallinson/forgery"
)

var store sessions.Store

type Sx struct {
  w http.ResponseWriter
  r *http.Request
  s *sessions.Session
}

func New(w http.ResponseWriter, r *http.Request) *Sx {
  session, _ := store.Get(r, "session")
  return &Sx{
    w: w,
    r: r,
    s: session,
  }
}

func Session(req *f.Request) *Sx {
  return req.Map["session"].(*Sx)
}

func (this *Sx) Save() {
  this.s.Save(this.r, this.w)
}

func (this *Sx) Set(key interface{}, value interface{}) {
  this.s.Values[key] = value
}

func (this *Sx) Get(key interface{}) interface{} {
  return this.s.Values[key]
}

func Middleware(secret string) (func(*stackr.Request, *stackr.Response, func())) {
  store = sessions.NewCookieStore([]byte(secret))
  return func(req *stackr.Request, res *stackr.Response, next func()) {
    req.Map["session"] = New(res.Writer, req.Request)
    next()
  }
}
