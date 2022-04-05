package ThreadComputation

import (
  "fmt"
  conv "github.com/cstockton/go-conv"
)


func TCFloatConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case float64:
    switch to_type {
    case String:
      return fmt.Sprintf("%v", e)
    case Int:
      ret, err := conv.Int64(int64(data.(float64)))
      if err == nil {
        return ret
      }
    case Float:
      ret, err := conv.Float64(float64(e))
      if err == nil {
        return ret
      }
    case Bool:
      if e == 1 {
        return true
      } else {
        return false
      }
    case Binary:
      res := MakeBinary(float64ToBytes(e, 8))
      res.Type = Float
      return res
    }
  }
  return nil
}



func init() {
  RegisterConvertCallback(Float, TCFloatConvert)
}
