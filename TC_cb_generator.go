package ThreadComputation

import (
  "fmt"
)



type TCGenFun func(interface{}) *TCIterator
type TCNextFun func(*TCIterator) interface{}


func GetGenCallback(v interface{}) TCGenFun {
  fn := fmt.Sprintf("gen.%v", TCType(v))
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCGenFun)
  }
  return nil
}

func RegisterGenCallback(gtype int, fun TCGenFun) {
  fname := fmt.Sprintf("gen.%v", gtype)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}

func GetNextCallback(v interface{}) TCNextFun {
  fn := fmt.Sprintf("next.%v", TCType(v))
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCNextFun)
  }
  return nil
}

func GetNextCallbackByType(ntype int) TCNextFun {
  fn := fmt.Sprintf("next.%v", ntype)
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCNextFun)
  }
  return nil
}

func RegisterNextCallback(gtype int, fun TCNextFun) {
  fname := fmt.Sprintf("next.%v", gtype)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}

func GetPrevCallback(v interface{}) TCNextFun {
  fn := fmt.Sprintf("prev.%v", TCType(v))
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCNextFun)
  }
  return nil
}

func GetPrevCallbackByType(ntype int) TCNextFun {
  fn := fmt.Sprintf("prev.%v", ntype)
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCNextFun)
  }
  return nil
}

func RegisterPrevCallback(gtype int, fun TCNextFun) {
  fname := fmt.Sprintf("prev.%v", gtype)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
