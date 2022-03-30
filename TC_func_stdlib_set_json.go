package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/Jeffail/gabs/v2"
)

func TCJsonSetString(x interface{}, y interface{}) {
  j1 := x.(*TCJson).J
  j2, err := gabs.ParseJSON([]byte(y.(string)))
  if err != nil {
    log.Debugf("Error parsing JSON: %v", err)
    return
  }
  j1.Merge(j2)
}

func TCJsonSetJson(x interface{}, y interface{}) {
  j1 := x.(*TCJson).J
  j2 := y.(*TCJson).J
  j1.Merge(j2)
}

func TCJsonSetDict(x interface{}, y interface{}) {
  j1 := x.(*TCJson).J
  j2 := y.(*TCDict).D
  for item := range j2.Items() {
    if TCisSimple(item.Value) {
      j1.Set(item.Value, item.Key.(string))
    }
  }
}

func init() {
  RegisterOperatorCmdCallback("set", Json, String, TCJsonSetString)
  RegisterOperatorCmdCallback("set", Json, Json, TCJsonSetJson)
  RegisterOperatorCmdCallback("set", Json, Dict, TCJsonSetDict)
}
