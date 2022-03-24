package ThreadComputation

import (
  "fmt"
)



type TCBlockFun func(*TCExecListener, string, string) interface{}



func GetBlockCallback(name string) TCBlockFun {
  fn := fmt.Sprintf("block.%v", name)
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCBlockFun)
  }
  return nil
}

func RegisterBlockCallback(name string, fun TCBlockFun) {
  fname := fmt.Sprintf("block.%v", name)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
