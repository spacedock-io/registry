package auth

import (
  "net/http"
  "encoding/gob"
  "regexp"
  "github.com/gorilla/sessions"
  "github.com/yawnt/registry.spacedock/context"
  "github.com/yawnt/registry.spacedock/repositories"
)

var parseToken = regexp.MustCompile(`^Token signature=(\w+),repository=(.*?),access=(\w+)$`)

type callback func(http.ResponseWriter, *http.Request)

func init() {
  gob.Register(&Token{})
}

type Token struct {
  Signature string
  Repo      repositories.Repository
  Access    string
}

func LoadCheckToken(s *sessions.Session, w http.ResponseWriter, r *http.Request) bool {
  if r.Header["Authorization"] == nil {
    return false
  }

  extracted := parseToken.FindStringSubmatch(r.Header["Authorization"][0])
  token := &Token{
    Signature: extracted[1],
    Repo: *repositories.NewRepo(extracted[2]),
    Access: extracted[3],
  }

  /*
   * Token Access must be compliant with the HTTP Method
   */
  if token.Access == "read" && r.Method != "GET" { return false }
  if token.Access == "write" && r.Method != "POST" && r.Method != "PUT" { return false }
  if token.Access == "delete" && r.Method != "DELETE" { return false }

  if token.Validate() {
    s.Values["token"] = token
    s.Save(r, w)
    return true
  }
  return false
}

func (t *Token) Header() string {
  return "Token signature=" + t.Signature + ",repository=" + t.Repo.String() + ",access="+ t.Access
}

func (t *Token) Validate() bool {
  client := &http.Client{}

  req, _ := http.NewRequest("GET", "https://index.docker.io/v1/repositories/" + t.Repo.String() + "/images", nil)
  req.Header.Add("Authorization", t.Header())

  resp, err := client.Do(req)
  if err == nil && resp.StatusCode == 200 {
    return true
  }
  return false
}

func Authenticated (w http.ResponseWriter, r *http.Request) bool {
  if r.URL.String() == "/" || r.URL.String() == "/v1/_ping" {
    return true
  }

  /*
   * Two types of auth are valid: Token or Session 
   */
  session, _ := context.Store.Get(r, "default")
  return session.Values["token"] != nil || LoadCheckToken(session, w, r)
}
