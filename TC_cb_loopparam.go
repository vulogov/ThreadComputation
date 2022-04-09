package ThreadComputation

import (
  "fmt"
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
)



type TCLoopParamFun func(*TCExecListener, interface{}, *deque.Deque) *deque.Deque


func GetLoopPCallback(x interface{}) TCLoopParamFun {
  x_type := TCType(x)
  fname := fmt.Sprintf("loopparam.%v", x_type)
  log.Debugf("Looking for loop param function for %T", x)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCLoopParamFun)
  }
  return nil
}

func RegisterLoopPCallback(x_type int, fun TCLoopParamFun) {
  fname := fmt.Sprintf("loopparam.%v", x_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
