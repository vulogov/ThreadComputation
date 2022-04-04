package ThreadComputation

import (

)

func TCLinesGet(x interface{}, y interface{}) interface{} {
  switch l := x.(type) {
  case *TCLines:
    switch n := y.(type) {
    case int64:
      if n < 0 || n > int64(l.Len()) {
        return nil
      }
      res := l.Q.At(int(n))
      if res == nil {
        return nil
      }
      return res
    }
  }
  return nil
}

func init() {
  RegisterOperatorCallback("get", Lines, Int, TCLinesGet)
}
