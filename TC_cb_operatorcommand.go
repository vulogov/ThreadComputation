package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
)



type TCOperatorCommand func(interface{},interface{})


func GetOperatorCmdCallback(name string, x interface{}, y interface{}) TCOperatorCommand {
  x_type := TCType(x)
  y_type := TCType(y)
  fname := fmt.Sprintf("opcmd.%v.%v.%v", name, x_type, y_type)
  log.Debugf("Looking for operator cmd for %v(%T %T)", name, x, y)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCOperatorCommand)
  }
  log.Debugf("No operator cmd for: %v(%T %T)", name, x, y)
  return nil
}

func RegisterOperatorCmdCallback(name string, x_type int, y_type int, fun TCOperatorCommand) {
  fname := fmt.Sprintf("opcmd.%v.%v.%v", name, x_type, y_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
