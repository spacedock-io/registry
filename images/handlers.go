package images

import(
  "net/http"
  "github.com/gorilla/mux"
)

func GetJson(w http.ResponseWriter, r *http.Request) {

}

func PutJson(w http.ResponseWriter, r *http.Request) {

}

func GetLayer(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  w.Write([]byte("hello " + params["id"]))
}

func PutLayer(w http.ResponseWriter, r *http.Request) {

}

func GetAncestry(w http.ResponseWriter, r *http.Request) {

}
