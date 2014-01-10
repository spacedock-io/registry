package main

import(
  "os"
  "github.com/codegangsta/cli"
  "github.com/codegangsta/martini"
  "github.com/ricallinson/forgery"
  "github.com/ricallinson/stackr"
  "github.com/spacedock-io/registry/router"
  "github.com/spacedock-io/registry/auth"
  "github.com/spacedock-io/registry/config"
  "github.com/Southern/logger"
)

const VERSION = "0.0.1"

func Secure(c *mux.Router) (func(http.ResponseWriter, *http.Request)) {
  return func(w http.ResponseWriter, r *http.Request) {

    if auth.Authenticated(w, r) {
      c.ServeHTTP(w, r);
    } else {
      w.WriteHeader(401);
    }
  }
}

func main() {
  server := f.CreateServer()
  server.Use(func (req *stackr.Request, res *stackr.Response, next func()) {
    defer next()

    res.Set("X-Docker-Registry-Version", VERSION)
    res.Set("X-Docker-Registry-Config", config.Env)
  })
  server.Use(auth.Middleware)

  app := cli.NewApp()

  app.Name = "Registry"
  app.Usage = "Run a standalone Docker registry"
  app.Version = context.VERSION
  app.Flags = []cli.Flag {
    cli.StringFlag{"port, p", "8080", "Port number", false},
    cli.StringFlag{"index, i", "false", "Index URL", false},
    cli.StringFlag{"env, e", "dev", "Environment", false},
  }

  app.Action = func(c *cli.Context) {
    env := c.String("env")
    if len(env) == 0 {
      env = "dev"
    }
    config.Global = config.Load(env)
    config.Logger = logger.New()

    context.Env = c.String("env")
    context.Port = c.String("port")
    context.Index = c.String("index")

    Routes(server)
    config.Logger.Log("Index listening on port " + fmt.Sprint(c.Int("port")))
    server.Listen(c.Int("port"))
  }

  app.Run(os.Args)
}
