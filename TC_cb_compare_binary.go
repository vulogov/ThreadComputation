package ThreadComputation

import (
  "bytes"
)





func TCCompareBinaryBinary(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCBinary:
    switch y.(type) {
    case *TCBinary:
      switch name {
      case "=":
        r := bytes.Compare(x.(*TCBinary).D, y.(*TCBinary).D)
        if r == 0 {
          return true
        }
      }
    }
  }
  return false
}

func TCCompareBinaryString(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCBinary:
    switch y.(type) {
    case string:
      switch name {
      case "=":
        r := bytes.Compare(x.(*TCBinary).D, []byte(y.(string)))
        if r == 0 {
          return true
        }
      }
    }
  }
  return false
}

func init() {
  RegisterCompareCallback(Binary, Binary, TCCompareBinaryBinary)
  RegisterCompareCallback(Binary, String, TCCompareBinaryString)
}
