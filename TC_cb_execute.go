package ThreadComputation

import (
  "fmt"
)



type TCExecuteFun func(*TCExecListener, interface{}) interface{}



func GetExecuteCallback(name string) TCBlockFun {
  fn := fmt.Sprintf("exec.%v", name)
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCBlockFun)
  }
  return nil
}

func RegisterExecuteCallback(etype int, fun TCExecuteFun) {
  fname := fmt.Sprintf("exec.%v", etype)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
