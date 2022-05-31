package ThreadComputation

import (

)



func TCListConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCList:
    switch to_type {
    case String:
      return e.String()
    case Numbers:
      return e.Numbers()
    }
  }
  return nil
}



func init() {
  RegisterConvertCallback(List, TCListConvert)
}
