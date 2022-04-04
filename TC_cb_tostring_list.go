package ThreadComputation

import (
  "fmt"
)



func TCConvertListToString(x interface{}) string {
  res := ""
  switch e := x.(type) {
  case *TCList:
    for i := 0; i < e.Len(); i++ {
      v := e.Q.At(i)
      switch v.(type) {
      case string:
        res += v.(string)
      default:
        fun := GetConverterCallback(v)
        if fun == nil {
          res += fmt.Sprintf("%v", v)
        } else {
          _res := fun(v, String)
          if _res != nil {
            res += _res.(string)
          } else {
            res += fmt.Sprintf("%v", v)
          }
        }
      }
    }
  }
  return res
}

func init() {
  RegisterToStringCallback(List, TCConvertListToString)
}
