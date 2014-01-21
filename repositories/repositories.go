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
    res.Send(err.Error(), 404)
    return
  }

  json, jsonErr := json.Marshal(tags)
  if jsonErr != nil {
    res.Send("Error sending data", 404)
    return
  }
  res.Send(json, 200)
}

func GetTag(req *f.Request, res *f.Response) {
  namespace := req.Params["namespace"]
  repo := req.Params["repo"]
  tag := req.Params["tag"]

  t, err := models.GetTag(namespace, repo, tag)
  if err != nil {
    res.Send(err.Error(), 404)
    return
  }

  json, jsonErr := json.Marshal(t)
  if jsonErr != nil {
    res.Send("Error sending data", 404)
    return
  }
  res.Send(json, 200)
}

func CreateTag(req *f.Request, res *f.Response) {
  namespace := req.Params["namespace"]
  repo := req.Params["repo"]
  tag := req.Params["tag"]
  json := req.Map["json"].(map[string]interface{})
  uuid, _ := json["uuid"].(string)

  err := models.CreateTag(namespace, repo, tag, uuid)
  if err != nil {
    res.Send(err.Error(), 500)
    return
  }

  res.Send("", 200)
}
