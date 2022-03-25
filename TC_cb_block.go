package ThreadComputation

import (
  "fmt"
  "strings"
)



type TCBlockFun func(*TCExecListener, string, string) interface{}



func GetBlockCallback(name string) TCBlockFun {
  name = strings.ToLower(name)
  fn := fmt.Sprintf("block.%v", name)
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCBlockFun)
  }
  return nil
}

func RegisterBlockCallback(name string, fun TCBlockFun) {
  name = strings.ToLower(name)
  fname := fmt.Sprintf("block.%v", name)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
