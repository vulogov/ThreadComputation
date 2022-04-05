package ThreadComputation

import (

)




func TCIOConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCIO:
    return e.String()
  }
  return nil
}



func init() {
  RegisterConvertCallback(IO, TCIOConvert)
}
