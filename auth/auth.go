package auth

import (
  "fmt"
  "net/http"
  "encoding/gob"
  "regexp"
  "github.com/gorilla/mux"
  "github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("MYSUPERSECRET"))
var parseToken = regexp.MustCompile(`^Token signature=(\w+),repository=(.*?),access=(\w+)$`)

type callback func(http.ResponseWriter, *http.Request)

type Token struct {
  Signature string
  Repo      string
  Access    string
}

func init() {
  gob.Register(&Token{})
}

func Secure(c *mux.Router) callback {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.URL.String() != "/" && r.URL.String() != "/v1/_ping" {
      session, _ := Store.Get(r, "default")
      if session.Values["token"] != nil || r.Header["Authorization"] != nil {
        if(session.Values["token"] == nil) {
          extracted := parseToken.FindStringSubmatch(r.Header["Authorization"][0])
          fmt.Println(extracted)
          session.Values["token"] = &Token{
            Signature: extracted[1],
            Repo: extracted[2],
            Access: extracted[3],
          }
          session.Save(r, w)
          fmt.Println(session.Values["token"].(*Token).Access)
        }
        c.ServeHTTP(w, r);
      } else {
        w.WriteHeader(401);
      }
    } else {
      c.ServeHTTP(w, r);
    }
  }
}
