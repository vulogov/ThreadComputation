package ThreadComputation

import (
  "fmt"
)



func TCConvertBinaryToString(x interface{}) string {
  res := ""
  switch x.(type) {
  case *TCBinary:
    return string(x.(*TCBinary).Raw())
  }
  return res
}

func init() {
  RegisterToStringCallback(Binary, TCConvertBinaryToString)
}
