package models

import (
  "github.com/spacedock-io/registry/db"
  "github.com/spacedock-io/registry/config"
)

func init() {
  config.Global = config.Load("dev")
  db.New(config.Global)
  db.DB.AutoMigrate(&Tag{})
}
