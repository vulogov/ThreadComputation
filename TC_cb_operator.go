package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
)



type TCGenericOperator func(interface{},interface{}) interface{}


func GetOperatorCallback(name string, x interface{}, y interface{}) TCGenericOperator {
  x_type := TCType(x)
  y_type := TCType(y)
  fname := fmt.Sprintf("op.%v.%v.%v", name, x_type, y_type)
  log.Debugf("Looking for operator for %v(%T %T)", name, x, y)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCGenericOperator)
  }
  log.Debugf("No operator for: %v(%T %T)", name, x, y)
  return nil
}

func RegisterOperatorCallback(name string, x_type int, y_type int, fun TCGenericOperator) {
  fname := fmt.Sprintf("op.%v.%v.%v", name, x_type, y_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
