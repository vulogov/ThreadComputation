package ThreadComputation

import (
  "fmt"
  "reflect"
  "unsafe"
)



type TCCloneFun func(interface{}) interface{}

func cloneString(s string) string {
  var b []byte
  h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
  h.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
  h.Len = len(s)
  h.Cap = len(s)
  return string(b)
}

func TCStringClone(data interface{}) interface{} {
  return cloneString(data.(string))
}

func TCSimpleClone(data interface{}) interface{} {
  return data
}

func GetCloneCallback(x interface{}) TCCloneFun {
  var fn string
  switch x.(type) {
  case int64:
    fn = fmt.Sprintf("clone.%v", Int)
  case float64:
    fn = fmt.Sprintf("clone.%v", Float)
  case string:
    fn = fmt.Sprintf("clone.%v", String)
  case bool:
    fn = fmt.Sprintf("clone.%v", Bool)
  default:
    fn = fmt.Sprintf("clone.%v", Any)
  }
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCCloneFun)
  }
  return nil
}

func RegisterCloneCallback(from_type int, fun TCCloneFun) {
  fname := fmt.Sprintf("clone.%v", from_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}

func init() {
  RegisterCloneCallback(Any, TCSimpleClone)
  RegisterCloneCallback(Int, TCSimpleClone)
  RegisterCloneCallback(Float, TCSimpleClone)
  RegisterCloneCallback(Bool, TCSimpleClone)
  RegisterCloneCallback(String, TCStringClone)
}
