package repositories

import (
  "strings"
  "encoding/json"
  "github.com/spacedock-io/registry/context"
  /* "github.com/spacedock-io/registry/images" */
  /* "github.com/garyburd/redigo/redis" */
)

type Repository struct {
  ns   string
  repo string
}

type JsonImg struct {
  Tag, id string
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

func (r *Repository) Put(data []byte) {
  var result []JsonImg
  err := json.Unmarshal(data, &result)

  /* context.Conn.Do("LPUSH", r.String() + ":_index_images" */
  /* conn.Do("H */
}
