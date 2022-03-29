package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
)



type TCGenericFunction func(interface{}) interface{}


func GetFunctionCallback(name string, x interface{}) TCGenericFunction {
  x_type := TCType(x)
  fname := fmt.Sprintf("fun.%v.%v", name, x_type)
  log.Debugf("Looking for function for %v(%T)", name, x)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCGenericFunction)
  }
  log.Debugf("No function for: %v(%T)", name, x)
  return nil
}

func RegisterFunctionCallback(name string, x_type int, fun TCGenericFunction) {
  fname := fmt.Sprintf("fun.%v.%v", name, x_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
