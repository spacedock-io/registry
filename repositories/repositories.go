package repositories

import (
  "encoding/json"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/registry/models"
)

func GetTags(req *f.Request, res *f.Response) {
  namespace := req.Params["namespace"]
  repo := req.Params["repo"]
  tags, err := models.GetTags(namespace, repo)
  if err != nil {
    res.Send(err.Error(), 400)
    return
  }

  json, jsonErr := json.Marshal(tags)
  if jsonErr != nil {
    res.Send("Error sending data", 400)
    return
  }
  res.Send(json, 200)
}
