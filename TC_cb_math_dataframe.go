package ThreadComputation

import (

)

func tcDataframeMathPair(name string, x *TCData, y *TCDict) *TCData {

  return x
}



func TCDataframeMath(name string, x interface{}, y interface{}) interface{} {
  switch d := x.(type) {
  case *TCData:
    switch v := y.(type) {
    case *TCDict:
      m := v.ToMap()
      d.D.Append(nil, m)
      return d
    }
  }
  return nil
}



func init() {
  RegisterMathCallback(Data, Dict, TCDataframeMath)
}
