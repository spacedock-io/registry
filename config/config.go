/*
Quick JSON configuration parser
*/
package config

import (
  "github.com/stretchr/objx"
  "github.com/Southern/logger"
  "io/ioutil"
  "log"
  "os"
  "path"
  "runtime"
)

var (
  Global objx.Map
  Logger *logger.Logger
  GoPath = os.Getenv("GOPATH")
  Dir string
)

func init() {
  _, file, _, ok := runtime.Caller(0)
  if ok {
    Dir = path.Dir(file)
  } else {
    if len(GoPath) > 0 {
      Dir = path.Join(
        GoPath,
        "src/github.com/spacedock-io",
        os.Args[0],
        "config",
      )
    }
  }
}

// Load takes an environment, loads the JSON file associated with the
// environment, and returns an instance of objx.Map for accessing the
// properties.
func Load(env string) (Global objx.Map) {
  // Get current directory
  pwd, err := os.Getwd()
  if err != nil {
    log.Fatalln(err)
  }

  // Try reading locally first
  data, localerr := ioutil.ReadFile(path.Join(pwd, "config",
    env + ".config.json"))
  if localerr != nil {
    // Try reading GOPATH next.
    if len(Dir) > 0 {
      data, err = ioutil.ReadFile(path.Join(
        Dir,
        env + ".config.json",
      ))
      if err != nil {
        log.Fatalln(err.Error() + "; " + localerr.Error())
      }
    } else {
      log.Fatalln(localerr.Error() +
        ", and $GOPATH is not defined to determine the config directory.")
    }
  }

  // Convert from JSON to objx.Map
  Global, err = objx.FromJSON(string(data))
  if err != nil {
    log.Fatalln(err)
  }

  return
}
