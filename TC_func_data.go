package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func IsSimpleData(data string) bool {
  fun := GetStringConverter()
  if fun == nil {
    return false
  }
  ar := [...]int{Int, Bool, Float, String}
  for _, x := range ar {
    e := fun(data, x)
    if e != nil {
      return true
    }
  }
  return false
}

func GetSimpleData(data string) interface{} {
  fun := GetStringConverter()
  if fun == nil {
    return nil
  }
  ar := [...]int{Int, Bool, Float, String}
  for _, x := range ar {
    e := fun(data, x)
    if e != nil {
      log.Debugf("Converting %v=%v, %v", data, e, x)
      return e
    }
  }
  return nil
}

func ApplyToFunction(l *TCExecListener, x interface{}, q *deque.Deque) interface{} {
  return x
}
