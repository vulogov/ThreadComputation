package ThreadComputation

import (

)




func TCLinesConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCLines:
    switch to_type {
    case Binary:
      return MakeBinary(e.String())
    case String:
      return e.String()
    }
  }
  return nil
}



func init() {
  RegisterConvertCallback(Lines, TCLinesConvert)
}
