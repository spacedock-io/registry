package images

import(
  "io"
  "encoding/json"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/registry/models"
  "github.com/spacedock-io/registry/cloudfiles"
)

func GetJson(req *f.Request, res *f.Response) {
  image, err := models.GetImage(req.Params["id"])
  if err != nil {
    res.Send(404)
    return
  }

  res.Set("X-Docker-Size", string(image.Size))
  res.Set("X-Docker-Checksum", image.Checksum)

  res.Send(image.Json)
}

func PutJson(req *f.Request, res *f.Response) {
  uuid := req.Params["id"]
  image, err := models.GetImage(uuid)
  if err != nil {
    if err != models.NotFoundErr {
      res.Send(err.Error(), 500)
      return
    }
    image = &models.Image{}
  }

  image.Uuid = uuid
  image.Json, err = json.Marshal(req.Map["json"])

  if err != nil {
    res.Send(500)
    return
  }

  err = image.Save()
  if err != nil {
    res.Send(err.Error(), 500)
    return
  }

  res.Send(200)
}

func GetLayer(req *f.Request, res *f.Response) {
  _, err := cloudfiles.Cloudfiles.ObjectGet(
    "spacedock", req.Params["id"], res.Response.Writer, true, nil)
  if err == nil {
    res.Send(200)
  } else { res.Send(500) }
}

func PutLayer(req *f.Request, res *f.Response) {
  obj, err := cloudfiles.Cloudfiles.ObjectCreate(
    "spacedock", req.Params["id"], true, "", "", nil)
  if err == nil {
    io.Copy(obj, req.Request.Request.Body)
    err = obj.Close()
    if err != nil {
      res.Send(500)
    } else { res.Send(200) }
  } else { res.Send(500) }
}

func GetAncestry(req *f.Request, res *f.Response) {
  image, err := models.GetImage(req.Params["id"])
  if err != nil {
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
