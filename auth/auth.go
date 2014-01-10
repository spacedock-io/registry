package auth

import (
  "regexp"
  "net/http"
  "github.com/ricallinson/forgery"
  /* "github.com/spacedock-io/registry/config" */
  "github.com/spacedock-io/registry/session"
)

var (
  tokenRegex = regexp.MustCompile(`^Token signature=(\w+),repository=(.*?),access=(\w+)$`)
)

type Token struct {
  Signature, Repo, Access string
}

func LoadCheckToken(req *f.Request) bool {
  header := req.Get("Authorization")
  if header != ""{
    return false
  }

  extracted := tokenRegex.FindStringSubmatch(header)
  token := &Token{
    Signature: extracted[1],
    Repo: extracted[2],
    Access: extracted[3],
  }

  /*
   * Token Access must be compliant with the HTTP Method
   */
  if token.Access == "read" && req.Method != "GET" { return false }
  if token.Access == "write" && req.Method != "POST" && req.Method != "PUT" { return false }
  if token.Access == "delete" && req.Method != "DELETE" { return false }

  if token.Validate() {
    sx.Session(req).Set("token", token)
    sx.Session(req).Save()
    return true
  }
  return false
}

func (t *Token) Header() string {
  return "Token signature=" + t.Signature + ",repository=" + t.Repo + ",access="+ t.Access
}

func (t *Token) Validate() bool {
  client := &http.Client{}

  req, _ := http.NewRequest("GET", "http://index.docker.io/v1/repositories/" + t.Repo + "/images", nil)
  req.Header.Add("Authorization", t.Header())

  resp, err := client.Do(req)
  if err == nil && resp.StatusCode == 200 {
    return true
  }
  return false
}

func Secure(route func(*f.Request, *f.Response)) func(*f.Request, *f.Response, func()) {
  /*
   * Two types of auth are valid: Token or Session 
   */
  return func(req *f.Request, res *f.Response, next func()) {
    defer next()
    if sx.Session(req).Get("token") != nil || LoadCheckToken(req) {
      route(req, res)
    } else {
      res.Send(401)
    }
  }
}
