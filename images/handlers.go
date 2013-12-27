package images

import(
  "binary"
  "net/http"
  "io/ioutil"
  "github.com/gorilla/mux"
  "launchpad.net/goamz/s3"
  "github.com/yawnt/registry.spacedock/context"
)

type Layer struct {
  checksum, json, size string
}

func GetJson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var layer Layer

  raw, _ := redis.Values(context.Conn.Do("GET", params["id"] + ":layer"))
  redis.ScanStruct(values, &layer)

  w.Header().Set("X-Docker-Size", layer.size)
  w.Header().Set("X-Docker-Checksum", layer.checksum)

  w.Write([]byte(layer.json))
}

func PutJson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  json, _ := ioutil.ReadAll(r.Body)

  context.Conn.Do("HSET", params["id"] + ":layer", "json", json)
}

func GetLayer(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  file, _ := context.FileStorage.Get(params["id"])

  w.Write(file)
}

func PutLayer(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  layer, _ := ioutil.ReadAll(r.Body)
  context.Conn.Do("HSET", params["id"] + ":layer", "size", binary.Size(layer))

  context.FileStorage.Put(params["id"], layer, "application/octet-stream", s3.Private)

  w.Write([]byte("hello " + params["id"]))
}

func GetAncestry(w http.ResponseWriter, r *http.Request) {

}
