package ThreadComputation

import (

)

func TCContextGet(x interface{}, y interface{}) interface{} {
  switch c := x.(type) {
  case *TCCtx:
    switch k := y.(type) {
    case string:
      if data, ok := c.C.Load(k); ok {
        return data
      }
    }
  }
  return MakeNone()
}

func init() {
  RegisterOperatorCallback("get", Context, String, TCContextGet)
}
