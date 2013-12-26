package images

import (
  "net/http"
  "github.com/gorilla/mux"
  /* "launchpad.net/goamz/s3" */
)

func GetLayer(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  w.Write([]byte("hello " + params["id"]))
}
