package main

import(
  "fmt"
  "os"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/yawnt/registry.spacedock/router"
  "github.com/yawnt/registry.spacedock/auth"
  "github.com/codegangsta/cli"
  "github.com/yawnt/registry.spacedock/context"
)

func Secure(c *mux.Router) (func(http.ResponseWriter, *http.Request)) {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("X-Docker-Registry-Version", context.VERSION)
    w.Header().Set("X-Docker-Registry-Config", context.VERSION)

    if auth.Authenticated(w, r) {
      c.ServeHTTP(w, r);
    } else {
      w.WriteHeader(401);
    }
  }
}

func main() {
  app := cli.NewApp()

  app.Name = "Registry"
  app.Usage = "Run a standalone Docker registry"
  app.Version = context.VERSION
  app.Flags = []cli.Flag {
    cli.StringFlag{"port, p", "8080", "Port number"},
    cli.StringFlag{"index, i", "false", "Index URL"},
    cli.StringFlag{"env, e", "dev", "Environment"},
  }

  app.Action = func(c *cli.Context) {
    context.Env = c.String("env")
    context.Port = c.String("port")

    fmt.Println("Registry listening on: http://127.0.0.1:" + context.Port)
    http.HandleFunc("/", Secure(router.Router))
    http.ListenAndServe(":" + context.Port, nil)
  }

  app.Run(os.Args)
}
