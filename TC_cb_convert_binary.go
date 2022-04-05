package ThreadComputation

import (

)


func TCBinaryConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCBinary:
    switch to_type {
    case String:
      return e.String()
    case Lines:
      return MakeLinesFromValue(e)
    }
  }
  return nil
}


func init() {
  RegisterConvertCallback(Binary, TCBinaryConvert)
}
