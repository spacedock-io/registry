package images

import(
  "fmt"
  "io"
  "io/ioutil"
  "encoding/json"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/registry/db"
  "github.com/spacedock-io/registry/models"
  "github.com/spacedock-io/registry/cloudfiles"
)

func GetJson(req *f.Request, res *f.Response) {
  image, err := models.GetImage(req.Params["id"])
  if q.Error != nil {
    res.Send(404)
    return
  }

  res.Set("X-Docker-Size", string(image.Size))
  res.Set("X-Docker-Checksum", image.Checksum)

  res.Send(image.Json)
}

func PutJson(req *f.Request, res *f.Response) {
  var image models.Image
  var err error

  image, err := models.GetImage(req.Params["id"])
  if q.Error != nil {
    if q.RecordNotFound() == false {
      res.Send(404)
      return
    }
  }

  image.Json, err = ioutil.ReadAll(req.Request.Request.Body)

  if err != nil {
    res.Send(500)
    return
  }

  fmt.Printf("image: %+v\n", image)
  err = i.Save()
  if err != nil {
    res.Send(err.Error(), 500)
    return
  }

  res.Send(200)
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
    "spacedock", req.Params["id"], true, "", "", nil)
  if err == nil {
    io.Copy(obj, req.Request.Request.Body)
    res.Send(200)
  } else { res.Send(500) }
}

func GetAncestry(req *f.Request, res *f.Response) {
  image, err := GetImage(req.Params["id"])
  if q.Error != nil {
    res.Send(404)
    return
  }

  data, err := json.Marshal(image.Ancestry)

  if err == nil {
    res.Send(data)
  } else { res.Send(err.Error(), 500) }
}

func PutChecksum(req *f.Request, res *f.Response) {
  uuid := req.Params["id"]
  /* *WTF* Docker?!
     HTTP API design 101: headers are *metadata*. The checksum should be passed
     as PUT body.
  */
  header := req.Header["X-Docker-Checksum"]
  if header == nil {
    res.Send("X-Docker-Checksum header is required", 400)
  }

  checksum := header[0]
  err := models.SetImageChecksum(uuid, checksum)
  if err != nil {
    res.Send(err.Error(), 500)
    return
  }

  res.Send(200)
}
