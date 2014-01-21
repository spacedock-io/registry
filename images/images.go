package images

import(
  "io"
  "io/ioutil"
  "encoding/json"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/registry/db"
  "github.com/spacedock-io/registry/models"
  "github.com/spacedock-io/registry/cloudfiles"
)

func GetJson(req *f.Request, res *f.Response) {
  var image models.Image
  db.DB.Where(&models.Image{Uuid: req.Params["id"]}).First(&image)

  res.Set("X-Docker-Size", string(image.Size))
  res.Set("X-Docker-Checksum", image.Checksum)

  res.Send(image.Json)
}

func PutJson(req *f.Request, res *f.Response) {
  var image models.Image
  var err error

  db.DB.Where(&models.Image{Uuid: req.Params["id"]}).First(&image)
  image.Json, err = ioutil.ReadAll(req.Request.Request.Body)

  if err == nil {
    db.DB.Save(&image)
  } else {
    res.Send(500)
  }
}

func GetLayer(req *f.Request, res *f.Response) {
  _, err := cloudfiles.Cloudfiles.ObjectGet(
    "default", req.Params["id"], res.Response.Writer, true, nil)
  if err == nil {
    res.Send(200)
  } else { res.Send(500) }
}

func PutLayer(req *f.Request, res *f.Response) {
  obj, err := cloudfiles.Cloudfiles.ObjectCreate(
    "default", req.Params["id"], true, "", "", nil)
  if err == nil {
    io.Copy(obj, req.Request.Request.Body)
    res.Send(200)
  } else { res.Send(500) }
}

func GetAncestry(req *f.Request, res *f.Response) {
  var image models.Image
  db.DB.First(&models.Image{Uuid: req.Params["id"]}).First(&image)

  data, err := json.Marshal(image.Ancestry)

  if err == nil {
    res.Send(data)
  } else { res.Send(500) }
}
