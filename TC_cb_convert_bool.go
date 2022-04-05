package ThreadComputation

import (
  "fmt"
)




func TCBoolConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case bool:
    switch to_type {
    case String:
      return fmt.Sprintf("%v", e)
    case Int:
      if e {
        return int64(1)
      } else {
        return int64(0)
      }
    case Float:
      if e {
        return float64(1)
      } else {
        return float64(0)
      }
    case Bool:
      return e
    case Binary:
      res := MakeBinary(boolToBytes(e))
      res.Type = Bool
      return res
    }
  }
  return nil
}


func init() {
  RegisterConvertCallback(Bool, TCBoolConvert)
}
