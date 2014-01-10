package router

import(
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/registry/images"
  "github.com/spacedock-io/registry/auth"
)

func Router(server *f.Server) {
  /* Home page */
  server.Get("/", func(req *f.Request, res *f.Response, next func()) {
    res.Send("docker-registry server")
  })

  /* Ping */
  server.Get("/v1/_ping", func(req *f.Request, res *f.Response, next func()) {
    res.Send("true")
  })

  /* Images Routes */
  server.Get("/v1/images/:id/ancestry", auth.Secure(images.GetAncestry))
  server.Get("/v1/images/:id/layer", auth.Secure(images.getLayer))
  server.Put("/v1/images/:id/layer", auth.Secure(images.putLayer))
  server.Get("/v1/images/:id/json", auth.Secure(images.getJson))
  server.Put("/v1/images/:id/json", auth.Secure(images.putJson))
}
