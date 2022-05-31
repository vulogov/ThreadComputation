package ThreadComputation

import (

)



func TCNumbersConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCNumbers:
    switch to_type {
    case String:
      return e.String()
    case List:
      return e.List()
    }
  }
  return nil
}


func init() {
  RegisterConvertCallback(Numbers, TCNumbersConvert)
}
