package repositories

import (
  "encoding/json"
  "fmt"
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

func GetTag(req *f.Request, res *f.Response) {
  namespace := req.Params["namespace"]
  repo := req.Params["repo"]
  tag := req.Params["tag"]

  t, err := models.GetTag(namespace, repo, tag)
  if err != nil {
    res.Send(err.Error(), 400)
    return
  }

  json, jsonErr := json.Marshal(t)
  if jsonErr != nil {
    res.Send("Error sending data", 400)
    return
  }
  res.Send(json, 200)
}

func CreateTag(req *f.Request, res *f.Response) {
  namespace := req.Params["namespace"]
  repo := req.Params["repo"]
  tag := req.Params["tag"]
  // json := req.Map["json"].(map[string]interface{})
  // uuid, _ := json["uuid"].(string)

  fmt.Println("Creating tag")
  fmt.Printf("ns: %s, repo: %s, tag: %s\n", namespace, repo, tag)

  // err := models.CreateTag(namespace, repo, tag, uuid)
  // if err != nil {
  //   res.Send(err.Error(), 400)
  //   return
  // }

  res.Send("", 200)
}
