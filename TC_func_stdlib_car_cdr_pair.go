package ThreadComputation

import (

)

func TCPairCar(v interface{}) interface{} {
  return v.(*TCPair).X
}

func TCPairCdr(v interface{}) interface{} {
  return v.(*TCPair).Y
}

func init() {
  RegisterFunctionCallback("car", Pair, TCPairCar)
  RegisterFunctionCallback("cdr", Pair, TCPairCdr)
}
