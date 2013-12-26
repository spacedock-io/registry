package auth

import (
  "net/http"
  "encoding/gob"
  "regexp"
  "github.com/gorilla/mux"
  "github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("MYSUPERSECRET"))
var parseToken = regexp.MustCompile(`^Token signature=(\w+),repository=(.*?),access=(\w+)$`)

type callback func(http.ResponseWriter, *http.Request)

func init() {
  gob.Register(&Token{})
}

type Token struct {
  Signature string
  Repo      string
  Access    string
}

func LoadCheckToken(s *sessions.Session, w http.ResponseWriter, r *http.Request) bool {
  if r.Header["Authorization"] == nil {
    return false
  } else {
    extracted := parseToken.FindStringSubmatch(r.Header["Authorization"][0])
    s.Values["token"] = &Token{
      Signature: extracted[1],
      Repo: extracted[2],
      Access: extracted[3],
    }
    s.Save(r, w)
    return s.Values["token"].(*Token).Validate()
  }
}

func (t *Token) Header() string {
  return "Token signature=" + t.Signature + ",repository=" + t.Repo + ",access="+ t.Access
}

func (t *Token) Validate() bool {
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "https://index.docker.io/v1/repositories/" + t.Repo + "/images", nil)
  req.Header.Add("Authorization", t.Header())
  resp, err := client.Do(req)
  if err != nil || resp.StatusCode != 200 {
    return false
  } else {
    return true
  }
}

func Secure(c *mux.Router) callback {
  return func(w http.ResponseWriter, r *http.Request) {
    /*
     * These two routes require no auth
     */
    if r.URL.String() != "/" && r.URL.String() != "/v1/_ping" {
      session, _ := Store.Get(r, "default")
      /*
       * Two types of auth are valid: Token or Session 
       */
      if session.Values["token"] != nil || LoadCheckToken(session, w, r) {
        c.ServeHTTP(w, r);
      } else {
        w.WriteHeader(401);
      }
    } else {
      c.ServeHTTP(w, r);
    }
  }
}
