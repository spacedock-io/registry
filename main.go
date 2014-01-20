package main

import(
  "fmt"
  "os"
  "github.com/codegangsta/cli"
  "github.com/ricallinson/forgery"
  "github.com/ricallinson/stackr"
  "github.com/spacedock-io/registry/db"
  "github.com/spacedock-io/registry/router"
  "github.com/spacedock-io/registry/config"
  "github.com/spacedock-io/registry/models"
  "github.com/spacedock-io/registry/session"
  "github.com/spacedock-io/registry/cloudfiles"
  "github.com/Southern/logger"
)

const VERSION = "0.0.1"

func main() {
  server := f.CreateServer()
  server.Use(func (req *stackr.Request, res *stackr.Response, next func()) {
    defer next()

    res.SetHeader("X-Docker-Registry-Version", VERSION)
    res.SetHeader("X-Docker-Registry-Config", "dev")
  })
  server.Use(sx.Middleware("SECRETVERYSECRET"))

  app := cli.NewApp()

  app.Name = "Registry"
  app.Usage = "Run a standalone Docker registry"
  app.Version = "0.0.1"
  app.Flags = []cli.Flag {
    cli.StringFlag{"port, p", "8080", "Port number"},
    cli.StringFlag{"index, i", "false", "Index URL"},
    cli.StringFlag{"env, e", "dev", "Environment"},
    cli.StringFlag{"config, c", "", "Configuration directory"},
  }

  app.Action = func(c *cli.Context) {
    env := c.String("env")
    dir := c.String("config")

    if len(env) == 0 {
      env = "dev"
    }
    if len(dir) > 0 {
      config.Dir = dir
    }

    config.Global = config.Load(env)
    config.Logger = logger.New()

    db.New(config.Global)
    db.DB.AutoMigrate(&models.Image{})
    db.DB.AutoMigrate(&models.Tag{})

    cloudfiles.New(config.Global)

    router.Routes(server)
    config.Logger.Log("Registry listening on port " + fmt.Sprint(c.Int("port")))
    server.Listen(c.Int("port"))
  }

  app.Run(os.Args)
}
