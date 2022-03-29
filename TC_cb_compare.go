package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
)



type TCCompareFun func(string, interface{}, interface{}) interface{}


func GetCompareCallback(x interface{}, y interface{}) TCCompareFun {
  x_type := TCType(x)
  y_type := TCType(y)
  fname := fmt.Sprintf("cmp.%v.%v", x_type, y_type)
  log.Debugf("Looking for compare for %v %v", x_type, y_type)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCCompareFun)
  }
  fname = fmt.Sprintf("cmp.%v.%v", x_type, Any)
  log.Debugf("Looking for compare for %v %v", x_type, Any)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCCompareFun)
  }
  fname = fmt.Sprintf("cmp.%v.%v", Any, y_type)
  log.Debugf("Looking for compare for %v %v", Any, y_type)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCCompareFun)
  }
  log.Debugf("No comparator for: %T %T", x, y)
  return nil
}

func RegisterCompareCallback(x_type int, y_type int, fun TCCompareFun) {
  fname := fmt.Sprintf("cmp.%v.%v", x_type, y_type)
  // fmt.Printf("Registering comparator: %v\n", fname)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
