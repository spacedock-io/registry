package router

import(
  "net/http"
  "github.com/gorilla/mux"
  "github.com/yawnt/registry.spacedock/images"
)

var Router = mux.NewRouter()

func init() {
  /* Home page */
  Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("docker-registry server"))
  }).Methods("GET")

  /* Ping */
  Router.HandleFunc("/v1/_ping", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("true"))
  }).Methods("GET")

  /* Images Routes */
  rImages := Router.PathPrefix("/v1/images/{id}").Subrouter()
  rImages.HandleFunc("/ancestry", images.GetAncestry).Methods("GET")
  rImages.HandleFunc("/layer", images.GetLayer).Methods("GET")
  rImages.HandleFunc("/layer", images.PutLayer).Methods("PUT")
  rImages.HandleFunc("/json", images.GetJson).Methods("GET")
  rImages.HandleFunc("/json", images.PutJson).Methods("PUT")
}
