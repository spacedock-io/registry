package main

import(
  "fmt"
  "os"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/yawnt/registry.spacedock/images"
  "github.com/yawnt/registry.spacedock/auth"
  "github.com/codegangsta/cli"
  "github.com/yawnt/registry.spacedock/context"
)

const VERSION = "0.0.1"

type Config struct {
  version, env string
}

func main() {
  app := cli.NewApp()

  app.Name = "Registry"
  app.Usage = "Run a standalone Docker registry"
  app.Version = VERSION
  app.Flags = []cli.Flag {
    cli.StringFlag{"port, p", "8080", "Port number"},
    cli.StringFlag{"index, i", "false", "Index URL"},
    cli.StringFlag{"env, e", "dev", "Environment"},
  }

  app.Action = func(c *cli.Context) {
    context.Set("version", VERSION)
    context.Set("env", c.String("env"))

    router := mux.NewRouter()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("docker-registry server"))
    }).Methods("GET")
    router.HandleFunc("/v1/_ping", func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("true"))
    }).Methods("GET")

    rImages := router.PathPrefix("/v1/images/{id}").Subrouter()
    rImages.HandleFunc("/ancestry", images.GetAncestry).Methods("GET")
    rImages.HandleFunc("/layer", images.GetLayer).Methods("GET")
    rImages.HandleFunc("/layer", images.PutLayer).Methods("PUT")
    rImages.HandleFunc("/json", images.GetJson).Methods("GET")
    rImages.HandleFunc("/json", images.PutJson).Methods("PUT")

    fmt.Println("Registry listening on: http://127.0.0.1:" + c.String("port"))

    http.HandleFunc("/", auth.Secure(router))
    http.ListenAndServe(":" + c.String("port"), nil)
  }

  app.Run(os.Args)
}
