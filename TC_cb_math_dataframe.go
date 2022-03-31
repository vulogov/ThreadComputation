package ThreadComputation

import (

)

func tcDataframeMathPair(name string, x *TCData, y *TCDict) *TCData {

  return x
}



func TCDataframeMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCData:
    switch y.(type) {
    case *TCPair:
      return tcDataframeMathPair(name, x.(*TCData), y.(*TCDict))
    }
  }
  return nil
}



func init() {
  RegisterMathCallback(Data, Dict, TCDataframeMath)
}
