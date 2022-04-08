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
    case Float:
      if e.Type == Float {
          return bytesToFloat64(e.D)
      }
    case Int:
      if e.Type == Int {
          return bytesToInt(e.D)
      }
    case Bool:
      if e.Type == Bool {
          return bytesToBool(e.D)
      }
    case Exported:
      return MakeExportedFromBinary(e)
    }
  }
  return nil
}


func init() {
  RegisterConvertCallback(Binary, TCBinaryConvert)
}
