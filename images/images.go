package images

import(
  "io"
  "encoding/json"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/registry/models"
  "github.com/spacedock-io/registry/cloudfiles"
  "strconv"
)

func GetJson(req *f.Request, res *f.Response) {
  image, err := models.GetImage(req.Params["id"])
  if err != nil {
    res.Send(404)
    return
  }

  res.Set("X-Docker-Size", strconv.Itoa(int(image.Size)))
  res.Set("X-Docker-Checksum", image.Checksum)

  res.Send(image.Json)
}

func PutJson(req *f.Request, res *f.Response) {
  var parent string
  _json := req.Map["json"].(map[string]interface{})
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

  val := _json["Size"]
  if val != nil {
    image.Size = int64(val.(float64))
  } else {
    image.Size = 0
  }
  _json["Size"] = int(image.Size)

  image.Json, err = json.Marshal(_json)

  if err != nil {
    res.Send(500)
    return
  }

  err = image.Save()

  p := _json["parent"]
  if p != nil {
    parent = p.(string)
  }
  e := updateAncestry(image, parent)

  if err != nil && e != nil {
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
    res.Send(500)
    return
  }

  res.Send(image.Ancestry)
}

func updateAncestry(image *models.Image, pId string) error {
  if pId == "" {
    ancestry := []string{image.Uuid}
    data, err := json.Marshal(ancestry)
    if err != nil {
      image.Ancestry = data
      return image.Save()
    }
    return err
  } else {
    var data []string
    pImage, err := models.GetImage(pId)
    json.Unmarshal(pImage.Ancestry, &data)

    newdata := make([]string, len(data)+1)
    newdata[0] = image.Uuid
    for i, elem := range data {
      newdata[i+1] = elem
    }

    if err != nil {
      return err
    }
    marshal, _ := json.Marshal(newdata)
    image.Ancestry = marshal

    return image.Save()
  }
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
