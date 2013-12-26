package repositories

import (
  "strings"
  /* "encoding/json" */
  "github.com/yawnt/registry.spacedock/context"
  /* "github.com/yawnt/registry.spacedock/images" */
  "github.com/garyburd/redigo/redis"
)

type Repository struct {
  ns   string
  repo string
}

func NewRepo(complete string) *Repository {
  parts := strings.Split(complete, "/")
  var repo, ns string
  if len(parts) < 2 {
    ns = "library"
    repo = parts[0]
  } else {
    ns = parts[0]
    repo = parts[1]
  }
  return &Repository{
    ns:   ns,
    repo: repo,
  }
}

func (r *Repository) String() string {
  return r.ns + "/" + r.repo
}

func (r *Repository) Put() {
  /* conn := context.Get("redis").(redis.Conn) */

  /* conn.Do("H */
}
