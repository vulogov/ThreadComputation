package ThreadComputation

import (

)


func TCExportedConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCExported:
    switch to_type {
    case String:
      return e.String()
    case Binary:
      return e.Export()
    }
  }
  return nil
}


func init() {
  RegisterConvertCallback(Exported, TCExportedConvert)
}
