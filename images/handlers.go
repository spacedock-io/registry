package images

import(
  "net/http"
  "io/ioutil"
  "github.com/gorilla/mux"
  "launchpad.net/goamz/s3"
  "github.com/yawnt/registry.spacedock/context"
)

func GetJson(w http.ResponseWriter, r *http.Request) {

}

func PutJson(w http.ResponseWriter, r *http.Request) {

}

func GetLayer(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  layer, _ := ioutil.ReadAll(r.Body)
  context.FileStorage.Put(params["id"], layer, "application/octet-stream", s3.Private)

  w.Write([]byte("hello " + params["id"]))
}

func PutLayer(w http.ResponseWriter, r *http.Request) {

}

func GetAncestry(w http.ResponseWriter, r *http.Request) {

}
