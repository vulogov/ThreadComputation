package ThreadComputation

import (

)

func TCNumbersGet(x interface{}, y interface{}) interface{} {
  switch l := x.(type) {
  case *TCNumbers:
    switch n := y.(type) {
    case int64:
      if n < 0 || n > int64(len(l.N)) {
        return nil
      }
      return l.N[int(n)]
    }
  }
  return nil
}

func init() {
  RegisterOperatorCallback("get", Numbers, Int, TCNumbersGet)
}
