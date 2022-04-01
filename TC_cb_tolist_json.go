package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)



func TCConvertJsonToList(x interface{}) *TCList {
  switch e := x.(type) {
  case *TCJson:
    res := new(TCList)
    for key, val := range e.J.ChildrenMap() {
      log.Debugf("json-to-list: %v=%v", key, val)
      p := new(TCPair)
      p.X = key
      p.Y = val
      res.Add(p)
    }
    return res
  }
  return nil
}

func init() {
  RegisterToListCallback(Json, TCConvertJsonToList)
}
